# TEST-Backend-Challenge

## Path "/test1/:testCase" 
    1 for testcase 1 
    expect response => 237 
    
    2 for testcase 2
    expect response => 7273
    
## Path "/test2/:encodeInput"  
    encodeInput for any hashed
    input => "LLRR="
    expect output => "210122"
    
    input => "==RLL" 
    expect output = "000210"

    input => "=LLRR"
    expect output = "221012"

    input => "RRL=R"
    expect output = "012001"

## Path "/beef/summary"
    expect => {
        "beef": {
            "ad": 46,
            "adipisicing": 42,
            "alcatra": 34,
            "aliqua": 34,
            "aliquip": 34,
            "andouille": 42
      }

## Path "/run-test"     
    for get all test result