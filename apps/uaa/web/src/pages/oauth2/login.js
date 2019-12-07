import React from 'react';
import ReactDOM from 'react-dom';
import Menu from 'components/Menu';

class LoginForm extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      username: "foo@bar.com",
      password: "foo@bar.com",
      challenge: window.loginData.challenge,
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
    // we need to submit form instead of ajax submit to follow redirects
  }

  render() {
    return (

        <form method="POST" onSubmit={this.handleSubmit}>
          <input type="hidden" name="challenge" value={this.state.challenge}/>
          <table>
            <tr>
              <td>
                <input id="username" name="username" type="email"
                       required="true"
                       value={this.state.username} onChange={this.setUsername}/>
              </td>
            </tr>
            <tr>
              <td>
                <input id="password" name="password" required="true"
                       value={this.state.password} onChange={this.setPassword}/>
              </td>
            </tr>
          </table>
          <input type="checkbox" id="remember" name="remember"
                 value="1"/><label htmlFor="remember">Remember
          me</label><br/><input type="submit" id="accept" value="Log in"/>
        </form>
    )
  }
}

ReactDOM.render(<Menu/>, document.getElementById('menu'));
ReactDOM.render(<LoginForm/>, document.getElementById('loginForm'));
