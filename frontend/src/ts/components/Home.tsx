
import {SetDefault, PropItems, StateTypes, SetStatus} from '..';
import {Login, Register, ComingSoon, Logo,mapDispatchToProps,mapStateToProps} from '.';

import React, {Component,SyntheticEvent} from 'react';
import { connect } from 'react-redux';

class Home extends Component<PropItems & StateTypes>{
  
  render(){
    let LState = this.props.Components.LoginPageState
    return(
      <div className="home_container">
        <img src={Logo} alt="log"/>

      {(LState) ? <Login/> : <Register/>}
      <div className='home_btn_container'>
        <button className={(LState)  ? "home_btn_active" : "home_btn_inactive"} onClick={()=>{this.props.Status({LoginPageState:true})}}>Login</button>
        <button className={(!LState) ? "home_btn_active" : "home_btn_inactive"} onClick={()=>{this.props.Status({LoginPageState:false})}}>Register</button> 
      </div>
       
      </div>
    )
  }
}
function Props(state:StateTypes){  
  return mapStateToProps(state);
}

function Dispatch(dispatch:any){  
  return mapDispatchToProps(dispatch);
}

export default connect(Props, Dispatch)(Home)



