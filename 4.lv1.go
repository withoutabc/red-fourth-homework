// 希望部分没实现，前面部分耗时太长了
package main

import (
	"fmt"
	"strings"
)

func main() {
	var skillNames, newSkill, newModel1, newModel2 string
	var num1, num2, num3, n int
	var y = 3
	var x = 1
	var z = 3
	slice := []string{"屎", "婊", "操", "他妈", "傻逼", "卷", "脑残", "智障", "内卷", "中国共产党", "二十大", "习近平", "疫情"}
	map3 := map[int]string{}
	map1 := map[int]string{
		1: "尝尝我的厉害吧!",
		2: "你的死期到了!",
		3: "接招吧!",
	}
	map2 := map[int]string{
		1: "龙卷风摧毁停车场!",
		2: "不要再卷辣!",
		3: "天崩!",
	}
	//以上为变量声明
	for {
		fmt.Printf("请选择以下操作\n1.使用技能\n2.自定义技能使用的文字模板\n3.添加自定义技能名字\n4.添加自定义的模板\n5.退出程序\n")
		fmt.Scan(&n)
		switch n {
		case 1:
		f1:
			for {
				fmt.Printf("请输入要使用的技能名\n")
				for x = 1; x <= y; x++ {
					fmt.Printf("%d.%v\n", x, map2[x])
				}
				//判断数字是否在所给列表
				for {
					fmt.Scan(&num2)
					_, ok2 := map2[num2]
					if ok2 == false {
						fmt.Printf("输入数字不在范围内！请重新输入\n")
						for x = 1; x <= y; x++ {
							fmt.Printf("%d.%v\n", x, map2[x])
						}
						continue
					}
					break
				}
				//EX1
				_, ok3 := map3[num2]
				//先看看是否已经添加过了模板
				if ok3 == true {
					fmt.Printf(map3[num2])
					//time.Sleep(1e9)
					fmt.Printf("\n请选择以下操作\n1.继续使用技能\n2.返回菜单\n")
				f3:
					for {
						fmt.Scan(&num3)
						switch num3 {
						case 1:
							break f3
						case 2:
							break f1
						default:
							fmt.Printf("输入数字不在范围内！请重新输入\n1.继续使用技能\n2.返回菜单\n")
							continue
						}
					}
				} else if ok3 == false {
					//没有模板就让用户选择模板
					fmt.Printf("请选择技能输入文字的模板\n")
					for x = 1; x <= z; x++ {
						fmt.Printf("%d.%v\n", x, map1[x])
					}
					for {
						fmt.Scan(&num1)
						_, ok1 := map1[num1]
						if ok1 == false {
							fmt.Printf("输入数字不在范围内！请重新输入\n")
							for x = 1; x <= y; x++ {
								fmt.Printf("%d.%v\n", z, map1[x])
							}
							continue
						}
						break
					}
					ReleaseSkill(skillNames, func(skillName string) {
						fmt.Printf("%v%v\n", map1[num1], map2[num2])
					})
					//time.Sleep(1e9)
					fmt.Printf("请选择以下操作\n1.继续使用技能\n2.返回菜单\n")
				f2:
					for {
						fmt.Scan(&num3)
						switch num3 {
						case 1:
							break f2
						case 2:
							break f1
						default:
							fmt.Printf("输入数字不在范围内！请重新输入！\n1.继续使用技能\n2.返回菜单\n")
							continue
						}
					}
				}
			}
		case 2: //设定模板
			fmt.Printf("请选择想要自定义文字模板的技能\n")
			for x = 1; x <= y; x++ {
				fmt.Printf("%d.%v\n", x, map2[x])
			}
			for {
				fmt.Scan(&num3)
				if num3 > y {
					fmt.Printf("输入数字不在范围内！请重新输入！\n")
					continue
				}
				break
			}
			fmt.Printf("请输入想要自定义的文字模板\n")
			fmt.Scan(&newModel1)
			map3[num3] = newModel1 + map2[num3]
			fmt.Println("模板设定成功！")
			//time.Sleep(1e9)
		case 3: //添加技能
			fmt.Printf("请输入想要添加的技能名\n")
			//检查步骤
		f4:
			for {
				fmt.Scan(&newSkill)
				for _, value1 := range slice {
					if strings.Index(newSkill, value1) != -1 {
						fmt.Printf("包含敏感词汇！请重新输入\n")
						continue f4
					}
				}
				//加入步骤
				y += 1
				map2[y] = newSkill
				fmt.Printf("技能添加成功！\n")
				break
			}
		case 4: //添加模板
			fmt.Printf("请输入你想添加的模板\n")
			//检查步骤
		f5:
			for {
				fmt.Scan(&newModel2)
				for _, value2 := range slice {
					if strings.Index(newModel2, value2) != -1 {
						fmt.Printf("包含敏感词汇！请重新输入\n")
						continue f5
					}
				}
				//加入步骤
				z += 1
				map1[z] = newModel2
				fmt.Printf("模板添加成功！\n")
				break
			}
		case 5:
			return
		}
	}
}

func ReleaseSkill(skillNames string, releaseSkillFunc func(string)) {
	releaseSkillFunc(skillNames)
}
