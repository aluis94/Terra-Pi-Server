#include json library
import json

#json string data
employee_string = '{"first_name": "Michael", "last_name": "Rodgers", "department": "Marketing"}'

#check data type with type() method
print(type(employee_string))

#convert string to  object
json_object = json.loads(employee_string)

#check new data type
print(json_object["first_name"])

b = 3>5
print(b)



def main():

    testCondition(50,51,"<")
    testCondition(50,51,">")
    testCondition(50,51,"==")
    testCondition(50,51,"!=")
    testCondition(50,51,"<=")
    testCondition(50,51,">=")

def testCondition(value,cvalue,operator):
    print("Testing the conditions")
    print("Values: ", value, cvalue)
    print("Operator: ",operator)

    isConditionMet = False
    if (operator == ">"):
        isConditionMet = value > cvalue
    elif (operator == "<"):
        isConditionMet = value < cvalue
    elif (operator == "=="):
        isConditionMet = value == cvalue
    elif (operator == "!="):
        isConditionMet = value != cvalue
    elif (operator == ">="):
        isConditionMet = value >= cvalue
    elif (operator == "<="):
        isConditionMet = value <= cvalue
    else:
        isConditionMet = False



    print("condition", isConditionMet)

main()
