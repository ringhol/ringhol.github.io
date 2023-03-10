
---
标题: 单一职责原则Single Responsibility Principle(SRP)
创建时间: 2023-03-08 11:13
修改时间: 2023-03-08 11:26
tags: 
---

![Pasted image 20220810212541](../attachments/Pasted%20image%2020220810212541.png)
现在有一个Telephone类作为本节的例子
![Pasted image 20220810215504](../attachments/Pasted%20image%2020220810215504.png)
若不遵从单一职责原则，我们的代码应该是这样的：
```cpp
/*
 *变化一:内部的变化，若Telephone类中四个函数任意一个需求改变，
 *      都需要修改Telephone类中的代码，不符合单一职责原则。
 *变化二: 外部的变化，如果Telephone要添加新的功能，也需要修改Telephone
 *      中的代码。
 */
class Telephone{
public:
	void dial(string phoneNumber){
		cout<<"拨打电话给："<<phoneNumber<<endl;
	}
	void hangUp(string phoneNumber){
		cout<<"挂断"<<phoneNumber<<"的电话"<<endl;
	}
	void sendMessage(string msg){
		cout<<"发送消息:"<<msg<<endl;
	}
	void receiveMessage(string msg){
		cout<<"收到消息:"<<msg<<endl;
	}
}
```
简单分析一下，由于这个类中干了四件事，引起这个类变化的因素有四个，当上面四个函数中有一个发生变化，都引起了这个类的变化。如果要添加功能，假如要添加开机和关机的功能，还是要修改这个类中的内容。
`重新设计的思路`：
	明确需求：我们希望有且只有一个引起Telephone类变化的原因（增加功能除外）。
	思路：给每个函数，都提炼成一个接口，然后编写接口的实现类，在Telephone类中只负责调用函数。
代码实现：
```cpp
//给每个函数都提炼成一个接口
class IDial{
public:
	virtual void dial(string phoneNumber) = 0;
};
class IHangUp{
public:
	virtual void hangUp(string phoneNumber) = 0;
};
class ISendMessage{
public:
	virtual void sendMessage(string msg) = 0;
};
class IReceiveMessage{
public:
	virtual void receiveMessage(string msg) = 0;
};
//编写接口的实现类
class Dial : public IDial{
public:
	void dial(string phoneNumber){
		cout<<"给"<<phoneNumber<<"拨打电话"<<endl;
	}
};
class HangUp : public IHangUp{
public:
	void hangUp(string phoneNumber){
		cout<<"挂断"<<phoneNumber<<"的电话"<<endl;
	}
};
class SendMessage : public ISendMessage{
public:
	void sendMessage(string msg){
		cout<<"发送消息:"<<msg<<endl;
	}
};
class ReceiveMessage : public IReceiveMessage{
public:
	void receiveMessage(string msg){
		cout<<"接收消息"<<msg<<endl;
	}
};
//在Telephone类中负责调用
class Telephone{
private:
	IDial* pDial;
	IHangUp* pHangUp;
	ISendMessage* pSend;
	IReceiveMessage* pReceive;
public:
	Telephone(IDial* dial,IHangUp* hangUp,ISendMessage* send,IReceiveMessage* recv):pDial(dial),pHangUp(hangUp),pSend(send),pReceive(recv){}
	void dial(string phoneNumber){
		pDial->dial(phoneNumber);
	}
	void hangUp(string phoneNumber){
		pHangUp->hangUp(phoneNumber);
	}
	void sendMessage(string msg){
		pSend->sendMessage(msg);
	}
	void receiveMessage(string msg){
		pReceive->receiveMessage(msg);
	}
};
```

符合单一职责原则的好处:
* `提高了代码的可读性`，提高了系统的可维护性。
* `降低类的复杂性`，一个模块只负责一个职责，提高系统的可扩展性和可维护性。
* `降低了变更所引起的风险`，变更时必然的，如果单一职责类做的好，当修改一个功能的时候可以显著的降低对另一个功能的影响。

课后练习：
将图中的关系用C++代码实现，要求符合单一职责原则
![Pasted image 20220810230638](../attachments/Pasted%20image%2020220810230638.png)
我的代码:
```cpp
#include<iostream>
using namespace std;

class IClassConstruct{
public:
    virtual void construct() = 0;
};

class IAttendenceStatistics{
public:
    virtual void attendenceStatistis() = 0;
};

class IPsychologicalCounseling{
public:
    virtual void psychologicalCounseling() = 0;
};

class IFeeConllection{
public:
    virtual void feeConllection() = 0;
};

class IClassManagent{
public:
    virtual void classManagent() = 0;
};

class ClassConstruct : public IClassConstruct{
public:
    void construct(){
        cout<<"班委建设工作完成"<<endl;
    }
};

class AttendenceStatistis : public IAttendenceStatistics{
public:
    void attendenceStatistis(){
        cout<<"出勤统计完成"<<endl;
    }
};

class PsychologicalCounseling : public IPsychologicalCounseling{
public:
    void psychologicalCounseling(){
        cout<<"心理辅导工作完成"<<endl;
    }
};

class FeeCollection : public IFeeConllection{
public:
    void feeConllection(){
        cout<<"费用催收工作完成"<<endl;
    }
};

class ClassManagement : public IClassManagent{
public:
    void classManagent(){
        cout<<"班级管理工作完成"<<endl;
    }
};

class ICounsellor{
public:
    virtual void classConstruct() = 0;
    virtual void attendenceStatistis() = 0;
    virtual void psychologicalCounseling() = 0;
    virtual void feeConllection() = 0;
    virtual void classManagent() = 0;
};

class Counsellor : public ICounsellor{
private:
    IClassConstruct* construct;
    IAttendenceStatistics* attend;
    IPsychologicalCounseling* psy;
    IFeeConllection* fee;
    IClassManagent* managent;
public:
    Counsellor(IClassConstruct* con,IAttendenceStatistics* att,IPsychologicalCounseling* psy,IFeeConllection* fee,IClassManagent*mg):construct(con),attend(att),psy(psy),fee(fee),managent(mg){}
    void classConstruct(){
        construct->construct();
    }
    void attendenceStatistis(){
        attend->attendenceStatistis();
    }
    void psychologicalCounseling(){
        psy->psychologicalCounseling();
    }
    void feeConllection(){
        fee->feeConllection();
    }
    void classManagent(){
        managent->classManagent();
    }
};

class IProfessionalGuidance{
public:
    virtual void professionalGuidance() = 0;
};

class IStudyGuidance{
public:
    virtual void studyGuidance() = 0;
};

class IScienceGuidance{
public:
    virtual void scienecGuidance() = 0;
};

class IJobGuidance{
public:
    virtual void jobGuidance() = 0;
};

class ProfessionalGuidance : public IProfessionalGuidance{
public:
    void professionalGuidance(){
        cout<<"专业指导完成"<<endl;
    }
};

class StudyGuidance : public IStudyGuidance{
public:
    void studyGuidance(){
        cout<<"学习指导完成"<<endl;
    }
};

class ScienceGuidance : public IScienceGuidance{
public:
    void scienecGuidance(){
        cout<<"科研指导完成"<<endl;
    }
};

class JobGuidance : public IJobGuidance{
public:
    void jobGuidance(){
        cout<<"就业指导完成"<<endl;
    }
};

class IMentor{
public:
    virtual void professionalGuidance() = 0;
    virtual void studyGuidance() = 0;
    virtual void scienceGuidance() = 0;
    virtual void jobGuidance() = 0;
};

class Mentor : public IMentor{
private:
    IProfessionalGuidance* profession;
    IStudyGuidance* study;
    IScienceGuidance* science;
    IJobGuidance* job;
public:
    Mentor(IProfessionalGuidance* ip,IStudyGuidance*stu,IScienceGuidance*sci,IJobGuidance*job):profession(ip),study(stu),science(sci),job(job){}
    void professionalGuidance(){
        profession->professionalGuidance();
    }
    void studyGuidance(){
        study->studyGuidance();
    }
    void scienceGuidance(){
        science->scienecGuidance();
    } 
    void jobGuidance(){
        job->jobGuidance();
    }
};

class IStudentWork{
public:
    virtual void classConstruct() = 0;
    virtual void attendanceStatistis() = 0;
    virtual void psychologicalGuidance() = 0;
    virtual void feeConllection() = 0;
    virtual void classManagent() = 0;
    virtual void professionalGuidance() = 0;
    virtual void studyGuidance() = 0;
    virtual void scienceGuidance() = 0;
    virtual void jobGuidance() = 0;
};

class StudentWork : public IStudentWork{
private:
    IMentor* mentor;
    ICounsellor* counsellor;
public:
    StudentWork(ICounsellor* counsellor,IMentor* mentor):mentor(mentor),counsellor(counsellor){}
    void classConstruct(){
        counsellor->classConstruct();
    }
    void attendanceStatistis(){
        counsellor->attendenceStatistis();
    }
    void psychologicalGuidance(){
        counsellor->psychologicalCounseling();
    }
    void feeConllection(){
        counsellor->feeConllection();
    }
    void classManagent(){
        counsellor->classManagent();
    }
    void professionalGuidance(){
        mentor->professionalGuidance();
    }
    void studyGuidance(){
        mentor->studyGuidance();
    }
    void scienceGuidance(){
        mentor->scienceGuidance();
    }
    void jobGuidance(){
        mentor->jobGuidance();
    }
};

void test01(){
    //实例化辅导员对象
    ClassConstruct con;
    AttendenceStatistis att;            
    PsychologicalCounseling psy;
    FeeCollection fee;
    ClassManagement mg;
    Counsellor counsellor(&con,&att,&psy,&fee,&mg);
    //实例化导师对象
    ProfessionalGuidance pro;
    StudyGuidance stu;
    ScienceGuidance sci;
    JobGuidance job;
    Mentor mentor(&pro,&stu,&sci,&job);    
    //实例化学生工作类
    StudentWork studentWork(&counsellor,&mentor);
    //测试studentWork中的功能
    studentWork.classConstruct();
    studentWork.attendanceStatistis();
    studentWork.psychologicalGuidance();
    studentWork.feeConllection();
    studentWork.classManagent();
    studentWork.professionalGuidance();
    studentWork.studyGuidance();
    studentWork.scienceGuidance();
    studentWork.jobGuidance();

}
int main(){
    test01();
    return 0;
}
```