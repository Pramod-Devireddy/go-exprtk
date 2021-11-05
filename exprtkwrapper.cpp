#include "exprtkwrapper.h"
#include "exprtk.hpp"
#include <iostream>
#include <vector>
#include <unordered_map>

struct DoubleValue{
    std::string varName;
    double value;
};

struct StringValue{
    std::string varName;
    std::string value;
};

struct VectorValue{
    std::string varName;
    std::vector<double> value;
};

struct  ExprtkStruct
{
    std::string exprString;
    std::vector<DoubleValue*> doubleValue;
    std::vector<StringValue*> stringValue;
    std::vector<VectorValue*> vectorValue;

    exprtk::symbol_table<double> symbolTable;
    exprtk::expression<double> expression;
    exprtk::parser<double> parser;
};

exprtkWrapper exprtkWrapperInit()
{
    ExprtkStruct* obj = new ExprtkStruct();
    return (void*)obj;
}


void setExpressionString(exprtkWrapper obj, char* expr)
{
    ExprtkStruct* exprtkStruct = (ExprtkStruct*)obj;
    exprtkStruct->exprString = std::string(expr);
}

const char* getExpressionString(exprtkWrapper obj)
{
    ExprtkStruct* exprtkStruct = (ExprtkStruct*)obj;
    return exprtkStruct->exprString.c_str();
}

int compileExpression(exprtkWrapper obj)
{
    ExprtkStruct* exprtkStruct = (ExprtkStruct*)obj;

    exprtkStruct->symbolTable.add_constants();
    exprtkStruct->expression.register_symbol_table(exprtkStruct->symbolTable);
    bool compiled = exprtkStruct->parser.compile(exprtkStruct->exprString, exprtkStruct->expression);
    if (compiled){
        return 1;
    }else{
        return 0;
    }
}

void setDoubleVariableValue(exprtkWrapper obj,  char* var, double val)
{
    ExprtkStruct* exprtkStruct = (ExprtkStruct*)obj;
    std::string varName = std::string(var);
    for (int i = 0; i < exprtkStruct->doubleValue.size();i++){
        if (varName.compare(exprtkStruct->doubleValue[i]->varName) == 0){
            exprtkStruct->doubleValue[i]->value = val;
            break;
        }
    }
}

void setStringVariableValue(exprtkWrapper obj,  char* var, char* val)
{
    ExprtkStruct* exprtkStruct = (ExprtkStruct*)obj;
    std::string varName = std::string(var);
    std::string value = std::string(val);
     for (int i = 0; i < exprtkStruct->stringValue.size();i++){
        if (varName.compare(exprtkStruct->stringValue[i]->varName) == 0){
            exprtkStruct->stringValue[i]->value = value;
            break;
        }
    }
}

void setVectorVariableValue(exprtkWrapper obj,  char* var, double val[], int len)
{
    ExprtkStruct* exprtkStruct = (ExprtkStruct*)obj;
    std::string varName = std::string(var);
    for (int i = 0; i < exprtkStruct->vectorValue.size();i++){
        if (varName.compare(exprtkStruct->vectorValue[i]->varName) == 0){
            exprtkStruct->vectorValue[i]->value.clear();
            for (int j = 0; j < len; j++){
                exprtkStruct->vectorValue[i]->value.push_back(double(val[j]));
            }
            break;
        }
    }
}
  

double getEvaluatedValue(exprtkWrapper obj)
{
    ExprtkStruct* exprtkStruct = (ExprtkStruct*)obj;
    return exprtkStruct->expression.value();
}

void addDoubleVariable(exprtkWrapper obj, char* var)
{
    ExprtkStruct* exprtkStruct = (ExprtkStruct*)obj;
    DoubleValue *d = new(DoubleValue);
    d->varName = std::string(var);
    d->value = 0;
    exprtkStruct->doubleValue.push_back(d);
    exprtkStruct->symbolTable.add_variable(d->varName,d->value);
}

void addStringVariable(exprtkWrapper obj, char* var)
{
    ExprtkStruct* exprtkStruct = (ExprtkStruct*)obj;
    StringValue *s = new(StringValue);
    s->varName = std::string(var);
    s->value = "";
    exprtkStruct->stringValue.push_back(s);
    exprtkStruct->symbolTable.add_stringvar(s->varName,s->value);
}

void addVectorVariable(exprtkWrapper obj, char* var)
{
    ExprtkStruct* exprtkStruct = (ExprtkStruct*)obj;
    VectorValue *v = new(VectorValue);
    v->varName = std::string(var);
    v->value.push_back(double(0));
    exprtkStruct->vectorValue.push_back(v);
    exprtkStruct->symbolTable.add_vector(v->varName,v->value);
}

void deleteExprtk(exprtkWrapper obj)
{
    delete (ExprtkStruct*)obj;
}