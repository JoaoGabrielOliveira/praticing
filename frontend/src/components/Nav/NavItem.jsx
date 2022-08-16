import React, {useState} from 'react';

    function NavItem(props){

    return (
        <button onClick={props.onClick}>
            <label>{props.label}</label>
        </button>
    )
}

export default NavItem;