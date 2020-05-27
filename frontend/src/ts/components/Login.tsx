import React from 'react';

const Login: React.FC = () => {
  return (

    <div>
    <form className="n_form" onSubmit={(e:any) =>{}} >
      <input autoComplete='username' id='user'  type="text" className="uname" placeholder={"Username"}/> 
      <input autoComplete='password' id='pass'  type="text" className="pwd"   placeholder={"Password"}/> 

      {/* <div className='home_btn_container'>
        <button className="home_btn" type='submit'>Login</button>
      </div> */}
    </form>
  <br />
  </div>
   
  );
}

export default Login;
// export default connect(mapStateToProps,mapDispatchToProps)(Login);

//------------------ HELPER FUNCTIONS ------------------


//------------------ PROP MAPPERS ------------------


// const ReactDefault: React.FC = () => {
//   return (
//     <div className="Login">
//       <header className="App-header">
//         <p>
//           Edit <code>src/App.tsx</code> and save to reload.
//         </p>
//         <a
//           className="App-link"
//           href="https://reactjs.org"
//           target="_blank"
//           rel="noopener noreferrer"
//         >
//           Learn React
//         </a>
//       </header>
//     </div>
//   );
// }