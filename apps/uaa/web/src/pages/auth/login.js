import React from 'react';
import ReactDOM from 'react-dom';
import Menu from 'components/Menu';

class LoginForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      username: "foo@bar.com",
      password: "foo@bar.com",
      redirect_url: window.loginData.redirectUrl,
    };
    this.handleSubmit = this.handleSubmit.bind(this);
    this.setUsername = this.setUsername.bind(this);
    this.setPassword = this.setPassword.bind(this);
  }

  setUsername(evt) {
    this.setState({username: evt.target.value})
  }

  setPassword(evt) {
    this.setState({password: evt.target.value})
  }

  handleSubmit(evt) {
  }

  render() {
    return (
        <form action="/sso" method="post" className="smart-green">
          <h1>SSO Login Form</h1>
          <label>
            <span>User Name :</span>
            <input id="username" name="username" type="email" required="true"
                   value={this.state.username} onChange={this.setUsername}/>
          </label>
          <label>
            <span>Password :</span>
            <input id="password" name="password" required="true"
                   value={this.state.password} onChange={this.setPassword}/>
          </label>
          <label>
            <span>&nbsp;</span>
            <input id="redirectUrl" type="hidden" name="redirectUrl"
                   value={this.state.redirectUrl}/>
          </label>
          <label>
            <span>&nbsp;</span>
            <input type="submit" className="button" value="Submit"/>
          </label>
        </form>
    )
  }
}

ReactDOM.render(<Menu/>, document.getElementById('menu'));
ReactDOM.render(<LoginForm/>, document.getElementById('loginForm'));
