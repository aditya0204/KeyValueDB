package test

import (
	"adiDB/server"
	"testing"
)

type KeyValues struct {
	Key string
	Value string
}





func BenchMark(i int, b *testing.B) {
	sd:=server.NewDB()
	example := []KeyValues{{Key:"x",Value:"1"},{Key:"x",Value:"1"},{Key:"x",Value:"1"},{Key:"x",Value:"1"},{Key:"x",Value:"1"},{Key:"x",Value:"1"},{Key:"x",Value:"1"},{Key:"x",Value:"1"},{Key:"x",Value:"1"},{Key:"x",Value:"1"},{Key:"x",Value:"1"},{Key:"x",Value:"1"},}
	for i:=0;i<b.N;i++{
		for _,el := range example{
			sd.Set(el.Key,el.Value)
		}
	}

}