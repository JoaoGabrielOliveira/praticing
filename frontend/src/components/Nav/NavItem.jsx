import React, {useState} from 'react';

    function NavItem(props){
    const [showOptions, setShowOptions] = useState(false)
    
    const handleShowOptions = () => {
        console.info("State is True");
        setShowOptions(true)
    }

    const handleCloseWhenBlur = () => {
        alert("Saiu de foco")
        setShowOptions(false)
    }

    let items;
    if (props.items?.length > 0){
        items = <ul>{props.items.map(element =>
                <li key={element.key}>{element}</li>
            )}</ul>;
    }


    return (
        <button onClick={props.onClick}>
            <label>{props.label}</label>
            {showOptions ? items : ""}
        </button>
    )
}

export default NavItem;