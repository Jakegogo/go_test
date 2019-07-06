package main

import (
	"debug/elf"
	"fmt"
	"log"
	"os"

	"github.com/bnagy/gapstone"
)

func main() {

	engine, err := gapstone.New(
		gapstone.CS_ARCH_X86,
		gapstone.CS_MODE_64,
	)
	if err != nil {
		log.Fatal(err)
	}

	path := os.Args[0]

	elfFile, err := elf.Open(path)
	if err != nil {
		log.Fatalf("error while opening ELF file %s: %+s", path, err.Error())
	}

	symbolTable, err := elfFile.Symbols()
	if err != nil {
		log.Fatalf("could not extract symbol table: %s", err.Error())
	}

	// extract the .text section
	textSection := elfFile.Section(".text")
	if textSection == nil {
		log.Fatal("No text section")
	}

	// extract the raw bytes from the .text section
	textSectionData, err := textSection.Data()
	if err != nil {
		log.Fatal(err)
	}

	// traverse through the symbol table
	for _, symbol := range symbolTable {

		if symbol.Size == 0 {
			continue
		}

		// skip over any symbols that aren't functinons/methods
		if symbol.Info != byte(2) && symbol.Info != byte(18) {
			continue
		}

		// calculate starting and ending index of the symbol within the text section
		symbolStartingIndex := symbol.Value - textSection.Addr
		symbolEndingIndex := symbolStartingIndex + symbol.Size

		// collect the bytes of the symbol
		symbolBytes := textSectionData[symbolStartingIndex:symbolEndingIndex]

		// disasemble the symbol
		instructions, err := engine.Disasm(symbolBytes, symbol.Value, 0)
		if err != nil {
			log.Fatalf("could not disasemble symbol: %s", err)
		}

		// print out each instruction that's part of this symbol
		fmt.Printf("\n\nSYMBOL %s\n", symbol.Name)
		for _, ins := range instructions {
			fmt.Printf("0x%x:\t%s\t\t%s\n", ins.Address, ins.Mnemonic, ins.OpStr)
		}
	}
}
