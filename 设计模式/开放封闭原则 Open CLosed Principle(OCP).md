
* 开放封闭原则是面向对象所有原则的核心。
* 两个关键点
	* 对功能扩展开放
	* 面向修改代码封闭
* 当需求改变时，在不改变软件实体源代码（类、接口、函数等）的前提下，通过拓展功能，使其满足新的需求。

本节案例：使用代码，描述不同需求的用户去银行办理不同的业务。
设计分析：
1. 用户:属性:记录不同类型的用户(存钱，取钱，转账......)
2. 银行柜员:帮助我们用户处理不同的需求
3. 银行业务系统:处理存钱、取钱、转账等需求的操作系统

普通的写法:
```cpp
#include<iostream>
using namespace std;
#include<string>

class BankClient{
private:
    string requirement;
public:
    void setRequirement(string requirement){
        this->requirement = requirement;
    }
    void setRequirement(const char * requirement){
        setRequirement(string(requirement));
    }
    string& getRequirement(){return requirement;}
};

//不符合单一职责原则
class BankSystem{
public:
    void deposite(){
        cout<<"处理存钱业务"<<endl;
    }
    void withdraw(){
        cout<<"处理取钱业务"<<endl;
    }
    void transfer(){
        cout<<"处理转账业务"<<endl;
    }
};

class BankStuff{
private:
BankSystem* system = new BankSystem();
public:
    void handleProcess(BankClient &client){
        string requirement = client.getRequirement();
        if(requirement == "存钱"){
            system->deposite();
        }else if(requirement == "取钱"){
            system->withdraw();
        }else if(requirement == "转账"){
            system->transfer();
        }else{
            cout<<"暂时无法处理您的需求~~~"<<endl;
        }
    }
};

void test01(){
    BankClient client;
    client.setRequirement("转账");
    BankStuff stuff;
    stuff.handleProcess(client);
}

int main(){
   test01();
   return 0; 
}
```
上述代码存在的`问题`：
* 不符合单一职责原则
* 想要添加功能时，不仅要修改BankSystem类，还要修改BankStuff类
* BankStuff类中的if过长

代码优化:
```cpp
#include<iostream>
using namespace std;

//BankSystem类中有大量的业务处理方法，将BankSystem类高度抽象
class IBankSystem{
public:
    virtual void bankProcess() = 0;
};

class DepositeClass : public IBankSystem{
public:
    void bankProcess(){
        cout<<"处理存款业务"<<endl;
    }
};

class WithDrawClass : public IBankSystem{
public:
    void bankProcess(){
        cout<<"处理取款业务"<<endl;
    }
};

class TransferClass : public IBankSystem{
public:
    void bankProcess(){
        cout<<"处理转账业务"<<endl;
    }
};

//将客户也进行抽象
//在过长的if语句中，引起if语句变化的原因是用户的需求不同
//不同的需求需要不同的业务处理对象
class IBankClient{
public:
    virtual IBankSystem* getBankSystem() = 0;
};

class DepositeClient : public IBankClient{
public:
    IBankSystem* getBankSystem(){
        return new DepositeClass();
    }
};

class WithDrawClient : public IBankClient{
public:
    IBankSystem* getBankSystem(){
        return new WithDrawClass();
    }
};

class TransferClient : public IBankClient{
    IBankSystem* getBankSystem(){
        return new TransferClass();
    }
};

class BankStuff{
private:
    IBankSystem* system;
public:
    void bankProcess(IBankClient* client){
        system = client->getBankSystem();
        system->bankProcess();
    }
};

void test01(){
    TransferClient client;
    BankStuff stuff;
    stuff.bankProcess(&client);
}

int main(){
    test01();
    return 0;
}
```
