import React from 'react';
import NavCollection from './NavCollection';
import NavItem from './NavItem';
import "./style.css";

/*
*/
//<NavItem label="Editar" />
//<NavItem label="Sobre" />

export default () => <nav>
    <NavCollection label="Arquivo">
        <NavItem label="Criar novo arquivo" onClick={() => alert("Novo")}/>
        <NavItem label="Abrir arquivo" onClick={() => alert("Abrir")}/>
        <NavItem label="Salvar arquivo" onClick={() => alert("Salvar")}/>
    </NavCollection>
</nav>
