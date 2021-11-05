#ifndef EXPRTKWRAPPER_HPP
#define EXPRTKWRAPPER_HPP

#ifdef __cplusplus
extern "C" {
#endif

typedef void* exprtkWrapper;

exprtkWrapper exprtkWrapperInit(void);

void setExpressionString(exprtkWrapper, char* expr);

const char* getExpressionString(exprtkWrapper);

void addDoubleVariable(exprtkWrapper, char* var);

void addStringVariable(exprtkWrapper, char* var);

void addVectorVariable(exprtkWrapper, char* var);

int compileExpression(exprtkWrapper);

void setDoubleVariableValue(exprtkWrapper, char* var, double val);

void setStringVariableValue(exprtkWrapper, char* var, char* val);

void setVectorVariableValue(exprtkWrapper, char* var, double val[], int len);

double getEvaluatedValue(exprtkWrapper);

void deleteExprtk(exprtkWrapper);


#ifdef __cplusplus
}
#endif

#endif // EXPRTKWRAPPER_HPP
