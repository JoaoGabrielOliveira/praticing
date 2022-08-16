import React, {useState} from "react"

export default (props) => {
    const [showOptions, setShowOptions] = useState(false)
    
    const handleShowOptions = () => {
        setShowOptions(true)
    }

    const handleBlur = () => {
        setShowOptions(false)
    }

    let items = <ul onBlur={handleBlur}>{props.children.map(element =>
        <li key={element.key}>{element}</li>
    )}</ul>;

    return <span className="menuDropdown">
        <button onClick={handleShowOptions}><label>{props.label}</label></button>
        {showOptions ? items : ""}
    </span>
}