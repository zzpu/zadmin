# BLL层全称是 Business Logic Layer
顾名思义,是业务层

换句话说,它是DAL(Data Access Layer,数据访问层)和UI(User Interface)层的连接桥梁.

既然称作业务层,必然有他的用处,不仅仅是一个中转的功能.
比如我要创建一个用户,可以用以下的逻辑表示:

namespace BLL
class 用户BLL
{
添加结果 AddUser(用户实体)
{
  if(!检查用户名是否合法(用户实体.用户名))return 用户名非法;
  if(!检查用户密码是否合法(用户实体.密码))return 密码非法;
  if(!DAL.检查用户是否存在(用户实体.用户名))return 用户名已经存在;
  int 新用户ID=DAL.添加用户记录(用户实体);
  if(新用户ID>0)return 用户添加成功;
  else reutrn 数据库访问出现错误!
}
}

但是在大部分没有严格要求的环境中,我们会习惯于把这些检查代码放在UI层,其实是不对的,从而造就了BLL层看起来就是一个中转的功能的错觉.
