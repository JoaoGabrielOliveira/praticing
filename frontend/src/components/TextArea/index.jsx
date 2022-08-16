import React, { useState } from 'react';

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