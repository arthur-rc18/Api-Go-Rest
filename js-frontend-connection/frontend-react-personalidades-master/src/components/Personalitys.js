import React, { Component } from 'react';
import axios from 'axios';
import '../components/Personalitys.css';

export default class Personalitys extends Component {
    state = {
        personalitys: []
    }

    componentDidMount() {
        axios.get('http://localhost:8000/api/personalitys')
            .then(res => {
                const personalitys = res.data;
                this.setState({ personalitys })
            })
    }

    render() {
        return (
            <div>
                {this.state.personalitys.map((p, id )=>
                    <div className="CardPersonalitys" key={id}>
                        <h3>{p.name}</h3>
                        <p>{p.history}</p>
                    </div>)}
            </div>
        );
    }
}
