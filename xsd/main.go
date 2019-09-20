package main

import (
	"fmt"
	"strings"
	"xslt-platform/common"
)

const(
	folder42 = `d:/go/code/xslt-platform/xsd/4.2`
	folder40 = `d:/go/code/xslt-platform/xsd/4.0`
	patern_attr=`<xs:attribute ref="(\w+)" use="required"/>`
)

func GetData(path string) map[string][]string {
	files := common.ListFile(path)
	result :=map[string][]string{}
	for _, f := range files {
		if strings.HasSuffix(f, ".xsd") {
			ar:= make([]string,0);
			s := common.ReadToString(f)
			items := common.GetData(s, patern_attr)
			for _, v := range items {
				ar=append(ar, v)
			}
			index := strings.LastIndex(f, "/")
			if index > 1 {
				f:=f[index+1:len(f)-4]
				result[f]=ar
			}
		}
	}
	return result
}
func main() {
	data42 := GetData(folder42)
	data40 := GetData(folder40)
	fmt.Println("s1000d 4.2 与 4.0 比较 新增和删除的 必填属性")
	for k,v42:= range data42 {
		fmt.Printf("============== %v start ===============\n",k)
		if v40,ok:=data40[k];ok{
			fmt.Println("4.2:"+k,len(v42))
			fmt.Println("4.0:"+k,len(v40))
			add := common.Difference(v42, v40)
			fmt.Printf("新增:%v\n",add)
			sub := common.Difference(v40,v42)
			fmt.Printf("删除:%v\n",sub)
		}else{
			fmt.Println("新增类型:"+k,v42)
		}
		fmt.Printf("============== %v end ===============\n",k)
	}

}
