/*
 * @Author: zhengzhuang
 * @Date: 2021-07-19 10:51:47
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-19 13:46:47
 * @Description: 接口
 * @FilePath: /01-study/面对对象/接口/main.go
 */
package main

import "fmt"

func main() {
	// realizationInterface()
	// oneToMany()
	manyToOne()
}

// 初级接口
type Sayer interface {
	say()
}
type dog struct{}
type cat struct{}

func (d dog) say() {
	fmt.Println("汪汪汪")
}
func (c cat) say() {
	fmt.Println("喵喵喵")
}

// 初级接口
func realizationInterface() {
	var x Sayer // 声明一个Sayer类型的变量x
	a := cat{}  // 实例化一个cat
	b := dog{}  // 实例化一个dog
	x = a       // 可以把cat实例直接赋值给x
	x.say()     // 喵喵喵
	x = b       // 可以把dog实例直接赋值给x
	x.say()     // 汪汪汪
}

// 一个类型实现多个接口
type Sayers interface {
	says()
}
type Mover interface {
	move()
}
type dogs struct {
	name string
}

func (d dogs) says() {
	fmt.Printf("%s会叫汪汪汪\n", d.name)
}

func (d dogs) move() {
	fmt.Printf("%s会动\n", d.name)
}

func oneToMany() {
	var x Sayers
	var y Mover

	var a = dogs{name: "旺财"}
	x = a
	y = a
	x.says()
	y.move()
}

// 多个类型实现同一接口
type Movers interface {
	moves()
}

type bird struct {
	name string
}

type fish struct {
	brand string
}

// bird类型实现Mover接口
func (d bird) moves() {
	fmt.Printf("%s 会跑\n", d.name)
}

// fish类型实现Mover接口
func (c fish) moves() {
	fmt.Printf("%s速度70迈\n", c.brand)
}
func manyToOne() {
	var x Movers
	var a = bird{name: "水"}
	var b = fish{brand: "鱼"}
	x = a
	x.moves()
	x = b
	x.moves()
}
