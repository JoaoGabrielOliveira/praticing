import React, { useState } from 'react';
import './style.css';

function TextArea() {
    const [text, setText] = useState("");

    /**
     * 
     * @param {KeyboardEvent} e 
     */
    function setTextOnState(e){
        setText(e.target.value);
    }

    return (
        <textarea id="text" onKeyDown={setTextOnState}></textarea>
    )
}

export default TextArea;