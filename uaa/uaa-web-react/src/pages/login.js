import React from 'react';
import ReactDOM from 'react-dom';
import Menu from 'components/Menu';

ReactDOM.render(<Menu/>, document.getElementById('menu'));
ReactDOM.render(<NameForm/>, document.getElementById('form'));


class NameForm extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            error: window.loginFormData.error,
            challenge: window.loginFormData.challenge,
            username: 'foo@bar.com',
            password: 'foo@bar.com'
        };

        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleChange(event) {
        this.setState({value: event.target.value});
    }

    handleSubmit(event) {
        alert('A name was submitted: ' + this.state.value);
        event.preventDefault();
    }

    render() {
        return (
            <form method={"POST"}>
                <input type="hidden" name="challenge" value={this.state.challenge}/>
                <table>
                    <tr>
                        <td><input id="username" name="username" placeholder="god"
                                   value={this.state.username}/></td>
                        <td>(it's "god")</td>
                    </tr>
                    <tr>
                        <td><input type="password" id="password" name="password"
                                   value={this.state.password}/></td>
                        <td>(it's "test123")</td>
                    </tr>
                </table>
                <input type="checkbox" id="remember" name="remember" value="1"/>
                <label htmlFor="remember">Remember
                    me</label><br/><input type="submit" id="accept" value="Log in"/>
            </form>
        );
    }
}
