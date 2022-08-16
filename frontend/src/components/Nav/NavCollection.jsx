import React, {useState} from "react"

export default (props) => {
    const [showOptions, setShowOptions] = useState(false)
    
    const handleToogleShow = () => {
        setShowOptions(!showOptions)
        alert("State is " + showOptions);
    }

    let items = <ul>{props.children.map(element =>
        <li key={element.key}>{element}</li>
    )}</ul>;

    //<button onClick={handleToogleShow}>{props.label}</button>
    return <span className="menuDropdown">
        <button onClick={handleToogleShow}><label>{props.label}</label></button>
        {showOptions ? items : ""}
    </span>
}