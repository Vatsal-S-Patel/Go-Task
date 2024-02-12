
// It returns inputField element
let inputField = document.getElementById('inputField');

// It returns all buttons element
let buttons = document.querySelectorAll('button');

// It's used for performing operations and storing temporary result and helps to update inputField 
let string = "";

// It keeps track that equal button is pressed or not
let isAnswerCalculated = false;

// It returns array of all button elements
let arr = Array.from(buttons);

// This will loop on all buttons and add click eventListeners to recognize that which button is pressed i.e. equal button, allClear button, clear button etc.
arr.forEach(button => {
    button.addEventListener('click', e => {

        // It store the HTMLElement that is pressed
        targetElement = e.target.innerHTML;

        // This statement evaluate if pressed button is equal button
        if (targetElement === '=') {
            calculateAnswer();
        }
        // This statement evaluate if pressed button is AC i.e. allClear
        else if (targetElement === 'AC') {
            allClear();
        }
        // This statement evaluate if pressed button is C i.e. clear or delete
        else if (targetElement === 'C') {
            clear();
        }
        // This statement evaluate if any number button or operator button is pressed
        else {
            takeInput(targetElement);
        } 
    });
});

// This function checks that passed element is Operator or not
function checkIsOperator(targetElement) {
    const operator = ['+','-','*','/','.'];

    // It'll return true if passed element is among operator array, otherwise return false
    return operator.includes(targetElement);
}

// This function update inputField's value by the current string's value 
function updateInputField() {
    inputField.value = string;    
}

// This function update by appending the string's value with passed element
function updateString(targetElement) {
    string += targetElement
}

// This function evaluate the inputField's value using eval() method and display it
function calculateAnswer() {
    // eval method evaluate the string expression and store result back into string
    string = eval(string);
    updateInputField();

    // As the equal button is pressed, the isAnswerCalculated variable toggles and become true, as it keeps track of status of equal button
    isAnswerCalculated = true;
}

// This function clears the inputField
function allClear() {
    // string's value become empty string and then inputField update by the string's value
    string = "";
    updateInputField();
}

// This functions delete the last element of inputField's value
function clear() {
    // It'll removes the last character of string's value and update string and update inputField value by string's value
    string = string.substring(0, string.length -1);
    updateInputField();
}

// This function take the input of Operator and Numbers also validates that operators are not coming consecutively, also update the string and inputField's value 
function takeInput(targetElement) {

    // isOperator is boolean variable that store the status of pressed button is operator or not
    let isOperator = false;

    // checkIsOperator method checks that pressed button is operator or not, if yes it become true, otherwise remain false
    isOperator = checkIsOperator(targetElement);

    // This statement will evaluate if the pressed button is not operator button and it's number button and answer has calculated
    if(!isOperator && isAnswerCalculated) {
        string = ""
        updateString(targetElement);
        updateInputField();
    } 
    // This statement checks that we pressed operator button previously and pressing operator button again second time, so it will replace that operator with current operator symbol
    else if (isOperator && inputField.value.length > 0 && checkIsOperator(inputField.value.slice(-1))) {
        clear()
        updateString(targetElement);
        updateInputField();
    } 
    // This statement evaluate if pressed button is operator button
    else {
        updateString(targetElement);
        updateInputField();
    }
    // After taking input that can be either number or operator, isAnswerCalculated will false
    isAnswerCalculated = false
}

// It adds the input functionality using KeyBoard 
document.addEventListener('keydown', event => {

    // This statement calculateAnswer if we press Enter key
    if (event.key === 'Enter') {
        calculateAnswer();
    }
    // This statement delete that last value if we press Enter key
    else if (event.key === 'Backspace') {
        clear();
    }
    // This statement clear the display if we press Delete key
    else if (event.key === 'Delete') {
        allClear();
    } 
    // This statement evaluate only if pressed key is number or operator, prevents the alphabets to input
    else if (parseInt(event.key) || event.key === '0' || checkIsOperator(event.key)){
        takeInput(event.key);
    } 
        
    // After taking input from keyboard it'll update the inputField
    updateInputField();
});