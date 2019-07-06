package main

import (
	"LogReplay-sdk-go/client/aspect/base"
	"LogReplay-sdk-go/client/utils"
	asp "github.com/Jakegogo/aspectgo/aspect"
	"net"
	"regexp"
)

// tcp连接初始化参数切面
// 回放时强制转成测试环境ip和端口
type NetConnAspect struct {
}

// Executed on compilation-time
func (a *NetConnAspect) Pointcut() asp.Pointcut {
	return asp.NewCallPointcutFromRegexp(`net\.DialTimeout|net\.Dial`)
}

// Executed ONLY on runtime
func (a *NetConnAspect) Advice(ctx asp.Context) []interface{} {
	args := ctx.Args()

	//回放
	if utils.IsReplay() {
		args[1] = []interface{}{"100.107.156.219:9092"}
	}

	return ctx.Call(args)
}


// 网络读切面
type NetConnReadAspect struct {
}

// Executed on compilation-time
func (a *NetConnReadAspect) Pointcut() asp.Pointcut {
	return asp.NewCallPointcutFromRegexp(`(\(net\.Conn\)\.Read|\(\*net\.conn\)\.Read)$`)
}

// Executed ONLY on runtime
func (a *NetConnReadAspect) Advice(ctx asp.Context) []interface{} {
	args := ctx.Args()

	//LogReplay开关关闭时，直接调用原有代码返回
	if !utils.IsLogReplayEnabled() {
		return ctx.Call(args)
	}

	//录制
	if utils.IsRecord() {
		res := ctx.Call(args)

		//缓存到redis
		//args[0] 本次读取的字节slice内容
		//res[0]  本次读取的字节slice长度
		//res[1]  err
		utils.RecordIntoRedis([]interface{}{args[0], res[0], res[1]})

		return res
	}

	//回放
	if utils.IsReplay() {
		var readBytes []byte
		var readLength int
		var err error
		var res = []interface{}{&readBytes, &readLength, &err}

		redisErrMaybe := utils.GetFromRedis(&res)
		if redisErrMaybe != nil {
			err = redisErrMaybe
		} else {
			dst := args[0].([]byte)
			copy(dst, readBytes)
		}

		return []interface{}{readLength, err}
	}

	return ctx.Call(args)
}

func (a *NetConnReadAspect) IsUseOuterType() bool {
	return true
}


// 网络写切面
type NetConnWriteAspect struct {
}

// Executed on compilation-time
func (a *NetConnWriteAspect) Pointcut() asp.Pointcut {
	return asp.NewCallPointcutFromRegexp(`(\(net\.Conn\)\.Write|\(\*net\.conn\)\.Write)$`)
}

// Executed ONLY on runtime
func (a *NetConnWriteAspect) Advice(ctx asp.Context) []interface{} {
	args := ctx.Args()

	//LogReplay开关关闭时，直接调用原有代码返回
	if !utils.IsLogReplayEnabled() {
		return ctx.Call(args)
	}

	//录制
	if utils.IsRecord() {
		res := ctx.Call(args)

		//缓存到redis
		//res[0]  本次写的字节长度
		//res[1]  err
		utils.RecordIntoRedis([]interface{}{res[0], res[1]})

		return res
	}

	//回放
	if utils.IsReplay() {
		var writeLength int
		var err error
		var res = []interface{}{&writeLength, &err}

		redisErrMaybe := utils.GetFromRedis(&res)
		if redisErrMaybe != nil {
			err = redisErrMaybe
		}

		return []interface{}{writeLength, err}
	}

	return ctx.Call(args)
}

func (a *NetConnWriteAspect) IsUseOuterType() bool {
	return true
}

type DialUdpAspect struct {
	base.BaseContextAspect
}

// Executed on compilation-time
func (a *DialUdpAspect) Pointcut() asp.Pointcut {
	pkg := regexp.QuoteMeta("net.DialUDP")
	return asp.NewCallPointcutFromRegexp(pkg)
}

func (a *DialUdpAspect) OnReplay() []interface{} {
	addr, _ := net.ResolveUDPAddr("udp", "100.107.156.219:9092")
	conn, err := net.DialUDP("udp", nil, addr)
	return []interface{}{conn, err}
}

type UDPConnWriteAspect struct {
	base.BaseContextAspect
}

// Executed on compilation-time
func (a *UDPConnWriteAspect) Pointcut() asp.Pointcut {
	return asp.NewCallPointcutFromRegexp(`\(*net\.UDPConn\)\.WriteTo(UDP)?$`)
}

func (a *UDPConnWriteAspect) OnRecord() error {
	utils.RecordIntoRedisWithKey(a.Res(), "udp_write")
	return nil
}

func (a *UDPConnWriteAspect) OnReplay() []interface{} {
	var ret int
	var err error
	res := []interface{}{&ret, &err}
	utils.GetFromRedisWithKey(&res, "udp_write")
	return []interface{}{ret, err}
}

type UDPConnReadFromAspect struct {
	base.BaseContextAspect
}

// Executed on compilation-time
func (a *UDPConnReadFromAspect) Pointcut() asp.Pointcut {
	return asp.NewCallPointcutFromRegexp(`\(*net\.UDPConn\)\.ReadFrom$`)
}

func (a *UDPConnReadFromAspect) OnRecord() error {
	utils.RecordIntoRedisWithKey(a.Res(), "udp_read")
	return nil
}

func (a *UDPConnReadFromAspect) OnReplay() []interface{} {
	var ret int
	var addr net.Addr
	var err error
	res := []interface{}{&ret, &addr, &err}
	utils.GetFromRedisWithKey(&res, "udp_read")
	return []interface{}{ret, addr, err}
}

type UDPConnReadFromUDPAspect struct {
	base.BaseContextAspect
}

// Executed on compilation-time
func (a *UDPConnReadFromUDPAspect) Pointcut() asp.Pointcut {
	return asp.NewCallPointcutFromRegexp(`\(*net\.UDPConn\)\.ReadFromUDP$`)
}

func (a *UDPConnReadFromUDPAspect) OnRecord() error {
	utils.RecordIntoRedisWithKey(a.Res(), "udp_read")
	return nil
}

func (a *UDPConnReadFromUDPAspect) OnReplay() []interface{} {
	var ret int
	var addr *net.UDPAddr
	var err error
	res := []interface{}{&ret, &addr, &err}
	utils.GetFromRedisWithKey(&res, "udp_read")
	return []interface{}{ret, addr, err}
}

// redis框架内部的切面，对redis返回包进行了录制
type NetConnReadStringAspect struct {
}

// Executed on compilation-time
func (a *NetConnReadStringAspect) Pointcut() asp.Pointcut {
	pointcut := regexp.QuoteMeta("(*bufio.Reader).ReadString")
	return asp.NewCallPointcutFromRegexp(pointcut)
}

// Executed ONLY on runtime
func (a *NetConnReadStringAspect) Advice(ctx asp.Context) []interface{} {
	args := ctx.Args()

	//LogReplay开关关闭时，直接调用原有代码返回
	if !utils.IsLogReplayEnabled() {
		return ctx.Call(args)
	}
	//录制
	if utils.IsRecord() {
		res := ctx.Call(args)

		//缓存到redis
		utils.RecordIntoRedis(res)

		return res
	}

	//回放
	if utils.IsReplay() {
		var str string
		var err error
		var res = []interface{}{&str, &err}
		redisErrMaybe := utils.GetFromRedis(&res)
		if redisErrMaybe != nil {
			err = redisErrMaybe
		}

		return []interface{}{str, err}
	}

	return ctx.Call(args)
}
