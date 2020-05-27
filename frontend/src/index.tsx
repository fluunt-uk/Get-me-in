import React from 'react';
import ReactDom from 'react-dom'
import { Provider } from "react-redux";
import {unregister,NavBar,store }  from './ts'; //Service Worker, Redux Store
import './css/index.css';

declare global { interface Window {store:any} }
window.store = store

ReactDom.render(<Provider store={store} > <NavBar/> </Provider>, document.getElementById("root"));

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
unregister();
