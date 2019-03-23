package main

// run as: go run typeface.go -f TestService.go -i IService -o IService.go -p main -s TestService
// gowrap: gowrap gen -p ./ -i IService -t fallback -o ./IServiceWrapper.go
type TestService struct {
}

func (s *TestService) Add(a, b int) (int, error) {
	return 0, nil
}
