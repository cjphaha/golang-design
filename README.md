# golang-design

## 单例模式

单例模式属于创建型模式，它提供了一种创建对象的最佳方式。该类负责创建自己的对象，同时确保只有单个对象被创建。

## 工厂模式

+ 简单工厂

  定义一个工厂类，根据参数的不同返回不同类的实例

+ 抽象工厂模式

  将一组具有同一主题的工厂封装起来，实际使用的时候需要实现这个抽象工厂。

+ DI(dependency injection) 依赖注入

  其实这个是 java 中常用的一种方式，go 原生并不支持 di，需要使用第三方库

  常用的主要有:

  + 反射实现的 dig https://github.com/uber-go/dig
  + 使用 go generate 实现的 wire https://github.com/google/wire
## 构造者模式

## 代理模式

> 代理类和被代理类实现同一个接口，代理类中持有被代理的对象

例如:
1. 火车票的代售票点。代售点就是代理，它拥有被大力对象的部门功能——售票功能

+ 静态代理
+ 动态代理
  java 中有字节码生成技术，但是 go 里面没有，可以用 go generate 来实现类似的功能

## 参考资料

+ https://lailin.xyz/post/factory.html
+ https://refactoringguru.cn/design-patterns/abstract-factory/go/example
+ https://learnku.com/articles/33707