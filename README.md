# Golang圣经
## 1.切片(Slice)
### 1.1切片的定义
切片的底层是数组,slice是对数组的封装，他描述的是一个数组的片段。
### 1.2切片的底层构造原理
slice实际上是一个struct,包含三个字段:底层数组,长度,容量。  
```
type slice struct {
	array unsafe.Pointer // 元素指针
	len   int // 长度 
	cap   int // 容量
}
```
### 1.3数组和切片的异同
相同:两者都可通过下标来进行随机访问。  
不同:  
1.数组是定长的,切片是不定长的。  
2.数组的长度是其类型的一部分,切片不是。  
3.切片支持动态扩容。
### 1.4切片的注意事项
底层数组是可以被多个 slice 同时指向的，因此对一个 slice 的元素进行操作是有可能影响到其他 slice 的。
### 1.5切片的扩容
#### 1.5.1切片扩容场景
使用 append 可以向 slice 追加元素，实际上是往底层数组添加元素。但是底层数组的长度是固定的，如果索引 len-1 所指向的元素已经是底层数组的最后一个元素，就没法再添加了。
这时，slice 会迁移到新的内存位置，新底层数组的长度也会增加，这样就可以放置新增的元素。
#### 1.5.2切片的扩容机制
Go1.18之前:  
当原 slice 容量小于 1024 的时候，新 slice 容量变成原来的 2 倍；原 slice 容量超过 1024，新 slice 容量变成原来的1.25倍。  

Go1.18之后:  
1.当原slice容量(oldcap)小于256的时候，新slice(newcap)容量为原来的2倍；  
2.原slice容量超过256，新slice容量newcap = oldcap+(oldcap+3*256)/4。  

注意：实际应用中并不是完全遵守以上规则,如果只看前半部分是相同的,但是后半部分,go会做内存对齐，内存对齐之后的新slice容量要大于等于前半部分生成的切片容量。
### 1.6切片作为函数参数
1.当 slice 作为函数参数时，就是一个普通的结构体。  
2.若直接传 slice, 实参slice并不会被函数中的操作改变；若传的是 slice 的指针，是会被函数中的操作改变原 slice 的。
3.不管传的是slice还是slice的指针,如果改变了 slice 底层数组的数据，会反应到实参slice 的底层数据。（比如s[1]=0这种操作）

## 2.Map
### 2.1Map的定义和作用
1.Map是由一组<key,value>对组成的数据结构,并且key只会出现一次。  
2.其作用设计一种数据结构用来维护一个集合的数据，并且可以同时对集合进行增删查改的操作。
### 2.2Map的底层实现
Golang中Map的底层实现是一个散列表,这个散列表中主要有两个结构体，一个叫hMap(表示map的结构体),一个叫bMap(bucket),其中bMap是在编译期间产生的。
### 2.3Map做为函数参数
1.当作为函数参数时，在函数参数内部对 map 的操作会影响 map 自身。  
2.因为创建map的时候,底层调用了makeMap这个函数,这个函数的返回值是一个指针。
### 2.4Map的扩容
Golang Map的扩容时机:   
1.装载因子是否大于6.5  
2.overflow bucket是否过多  

扩容方法:  
1.双倍扩容:扩容采取了一种称为“渐进式”地方式，原有的key 并不会一次性搬迁完毕，每次最多只会搬迁 2个bucket。  
2.等量扩容：重新排列，极端情况下，重新排列也解决不了，map 成了链表，性能大大降低，此时哈希种子 hash0 的设置，可以降低此类极端场景的发生。  

装载因子(load factor):用于衡量当前哈希表中空间占用率的核心指标，也就是每个 bucket 桶存储的平均元素个数。
### 2.5为什么map遍历结果是无序
1.因为map中key是无序的    
2.如果map发生扩容,会发生 key 的搬迁,原来落在同一个 bucket 中的 key在搬迁后,有些 key 就要远走高飞了。但是遍历的过程是按顺序遍历 bucket,同时按顺序遍历 bucket 中的 key。
搬迁后，key 的位置发生了重大的变化,有些 key 飞上高枝，有些 key 则原地不动。这样,遍历 map 的结果就不可能按原来的顺序了。
3.Golang中遍历map并不是固定的从0号bucket开始遍历的，是随机选择bucket开始遍历,并且也是从这个bucket的一个随机序号的cell开始遍历。
### 2.6如何让map遍历结果有序
1.把key拿出来放入一个slice中,把slice进行排序  
2.利用官方库里的 list(链表) 封装一个结构体,实现一个有序的 K-V 存储结构,在里面维护一个 keys 的 list。
### 2.7map可以边遍历边删除吗
1.如果发生在多个协程同时读写一个map的情况下,如果被检测到,直接panic。  
2.如果在同一个协程内,理论上可以,但是遍历的结果可能不同。
### 2.8如何比较两个map是否相等
不能直接比较,只能遍历map的每个元素,比较元素是否深度相等。  
深度相等的条件:  
1.都为nil  
2.非空,长度相等,指向同一个map实体对象。  
3.相应的key指向的value深度相等。

## 3.接口(Interface)
### 3.1值接收者和指针接收者的区别
1.指针接收者方法会改变struct内部的数据。  
2.实现了接收者是值类型的方法，相当于自动实现了接收者是指针类型的方法；而实现了接收者是指针类型的方法，不会自动生成对应接收者是值类型的方法。
### 3.2iface和eface的区别
iface 和 eface 都是 Go 中描述接口的底层结构体，区别在于 iface 描述的接口包含方法，而 eface 则是不包含任何方法的空接口：interface{}。  
### 3.3类型转换和断言的区别
1. 类型转换、类型断言本质都是把一个类型转换成另外一个类型。不同之处在于，类型断言是对接口变量进行的操作。  
2. 类型断言代码实例
```
type Student struct {
	Name string
	Age int
}
func main() {
	var i interface{} = new(Student)
	s, ok := i.(Student) //<目标类型的值>，<布尔参数> := <表达式>.( 目标类型 )
	if ok {
		fmt.Println(s)
	}
}
```
### 3.4多态
多态是一种运行期的行为，它有以下几个特点：  
1.一种类型具有多种类型的能力。   
2.允许不同的对象对同一消息做出灵活的反应。    
3.以一种通用的方式对待个使用的对象。   
4.非动态语言必须通过继承和接口的方式来实现。   

## 4.Goroutine
### 4.1什么是Goroutine
Goroutine是Go中协程的实现,是Go中的基本执行单元,是一个与其它goroutine并行运行在同一地址空间的函数或方法。  
### 4.2Goroutine和线程的区别
1.Goroutine(2KB)的内存占用比线程(1MB)小很多。  
2.Goroutine创建和销毁的消耗比线程小,因为线程是内核态,Goroutine是用户态。  
3.Goroutine的切换成本比线程小,因为Goroutine切换的时候只需要三个寄存器。
### 4.3GMP是什么
GMP是Go调度的三个核心组件  
G:代表一个goroutine,它包含:表示goroutine栈的一些字段,指示当前 goroutine 的状态,指示当前运行到的指令地址,也就是 PC 值。    
M:对内核级线程的封装,数量对应真实的CPU数。  
P:即为G和M的调度对象,用来调度G和M之间的关联关系,它维护一个处于Runnable状态的g队列,需要获得p才能运行g。  
### 4.4为什么要有P
调度器把G都分配到M上,不同的G在不同的M并发运行时,都需要向系统申请资源,比如堆栈内存等。因为资源是全局的,就会因为资源竞争照成很多性能损耗。为了解决这一的问题go从 1.1 版本引入,在运行时系统的时候加入p对象,让P去管理这个G对象,M想要运行G,必须绑定 P,才能运行P所管理的G。
### 4.5Goroutine的调度时机有哪些
1. 使用关键字Go: go创建一个新的goroutine,Go scheduler会考虑调度。  
2. GC:由于进行GC的goroutine也需要在M上运行,因此肯定会发生调度。当然,Go scheduler还会做很多其他的调度,例如调度不涉及堆访问的 goroutine 来运行。GC不管栈上的内存,只会回收堆上的内存。    
3. 系统调用: 当goroutine进行系统调用时,会阻塞 M,所以它会被调度走,同时一个新的goroutine会被调度上来。
4. 内寸同步访问: atomic,mutex,channel操作等会使goroutine阻塞,因此会被调度走。等条件满足后（例如其他 goroutine 解锁了）还会被调度上来继续运行。
### 4.6GMP调度流程
1. 每个P有个局部队列,局部队列保存待执行的goroutine,当M绑定的P的的局部队列已经满了之后就会把goroutine放到全局队列。  
2. 每个P和一个M绑定,M是真正的执行P中 goroutine的实体,M从绑定的P中的局部队列获取G来执行。  
3. 当M绑定的P的局部队列为空时,M会从全局队列获取到本地队列来执行G),当从全局队列中没有获取到可执行的G时候,M会从其他P的局部队列中偷取G来执行,这种从其他P偷的方式称为work stealing。  
4. 当G因系统调用(syscall)阻塞时会阻塞M,此时P会和M解绑即hand off,并寻找新的闲置的M,若没有idle的M就会新建一个 M。  
5. 当G因channel或者network I/O阻塞时,不会阻塞M,M会寻找其他就绪的G。当阻塞的G恢复后会重新进入就绪状态,进入 P 队列等待执行。
### 4.7Goroutine的调度方式
1.在Go1.14版本之前Goroutine是协作式的抢占式调度(程序只能依靠Goroutine自己交出CPU资源才能触发调度),存在以下问题。   
1.1 某些Goroutine可以长时间占用线程,造成其它Goroutine的饥饿。  
1.2 垃圾回收需要暂停整个程序(Stop-the-world,STW),最长可能需要几分钟的时间,导致整个程序无法工作。  
2.在1.14版本之后Go采用了基于信号的抢占式调度(异步抢占)。  
2.1 M注册一个 SIGURG信号的处理函数：sighandler。  
 sysmon 线程检测到执行时间过长的 goroutine、GCstw时,会向相应的M(或者说线程,每个线程对应一个M)发送SIGURG信号。   
 收到信号后,内核执行sighandler函数,通过pushCall插入asyncPreempt函数调用。  
2.2回到当前goroutine执行asyncPreempt函数,通过mcall切到g0栈执行gopreempt_m。  
2.3将当前goroutine插入到全局可运行队列,M则继续寻找其他goroutine来运行。  
2.4被抢占的goroutine再次调度过来执行时,会继续原来的执行流。  
### 4.8如何关闭一个Goroutine










