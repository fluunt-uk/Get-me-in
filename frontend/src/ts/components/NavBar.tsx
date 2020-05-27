import { connect } from "react-redux";
import { Component } from "react";

import { BrowserRouter as Router, Route,NavLink  } from "react-router-dom"

import React from "react";
import '../../css/App.css';

import { Home,Ads,mapStateToProps} from '.'
import { StateTypes } from "..";
class NavBar extends Component<StateTypes>{
    render(){
       return ( 
            <Router>
                <div>
                    <nav>
                        <div className='nav_global'>
                            <div  id='nav_bar_item'>
                                <NavLink  to='/' exact activeClassName='active'> Home </NavLink >
                            </div>
                            <div id='nav_bar_item'>
                                <NavLink  to='/register' exact activeClassName='active' > Page 1  </NavLink >
                            </div>
                        </div>
                    </nav>
                    <Route path ='/'        exact component={Home} /> 
                    <Route path='/ads' exact component={Ads} />
                </div>
            </Router>
        )   
    }
}

export default connect(mapStateToProps)(NavBar)