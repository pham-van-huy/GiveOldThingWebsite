import React from 'react'

class CreatePost extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            title: "",
            description: "",
            city: "",
            category: "",
            images: []
        };
    }
    
    componentDidMount() {
        fetch('/api/cities').then((results) => {
            console.log('results')
        })
    }
    render() {
        return (
            <form>
                <div className="form-row">
                    <div class="form-group col-md-6">
                        <label for="title">Title</label>
                        <input type="text" name="title" className="form-control" id="title" placeholder="Tieu de" />
                    </div>
                </div>
                <div className="form-group">
                    <label for="description">Description</label>
                    <textarea type="text" name="description" className="form-control" id="description" placeholder="" />
                </div>
                <div className="form-row">
                    
                </div>
                <div class="form-group">
                    
                </div>
                <button type="submit" class="btn btn-primary">Sign in</button>
            </form>
        );
    }
}

export default CreatePost
