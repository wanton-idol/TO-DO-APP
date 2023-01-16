import React, { Component } from "react";
import axios from "axios";
import { Card, Header, Form, Input, Icon, Button } from "semantic-ui-react";

let endpoint = "http://localhost:3000";

class ToDo extends Component {
  constructor(props) {
    super(props);

    this.state = {
      task: "",
      items: [],
    };
  }

  componentDidMount() {
    this.getTask();
  }

  onChange = (event) => {
    this.setState({
      [event.target.name]: event.target.value,
    });
  };

  onSubmit = () => {
    let { task } = this.state;
    if (task) {
      axios
        .post(
          endpoint + "/api/task",
          {
            task,
          },
          {
            headers: {
              "Content-Type": "application/x-www-form-urlencoded",
            },
          }
        )
        .then((res) => {
          this.getTask();
          this.setState({
            task: "",
          });
          console.log(res);
        });
    }
  };

  getTask = () => {
    axios.get(endpoint + "/api/task").then((res) => {
      if (res.data) {
        this.setState({
          items: res.data.map((item) => {
            let color = "orange";
            let style = {
              wordWrap: "break-word",
            };

            if (item.status) {
              color = "green";
              style["textDecorationLine"] = "line-through";
            }

            return (
              <Card key={item._id} color={color} fluid>
                <Card.Content
                  style={{
                    backgroundColor: "beige",
                  }}
                >
                  <Card.Header textAlign="left">
                    <div style={style}>{item.task}</div>
                  </Card.Header>

                  <Card.Meta textAlign="right">
                    <Button
                      color="green"
                      onClick={() => this.updateTask(item._id)}
                    >
                      <Icon name="check" /> Done
                    </Button>
                    <Button
                      color="yellow"
                      onClick={() => this.undoTask(item._id)}
                    >
                      <Icon name="undo" /> Undo
                    </Button>
                    <Button
                      color="red"
                      onClick={() => this.deleteTask(item._id)}
                    >
                      <Icon name="delete" /> Delete
                    </Button>
                  </Card.Meta>
                </Card.Content>
              </Card>
            );
          }),
        });
      } else {
        this.setState({
          items: [],
        });
      }
    });
  };

  updateTask = (id) => {
    axios
      .put(endpoint + "/api/task/" + id, {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
        },
      })
      .then((res) => {
        console.log(res);
        this.getTask();
      });
  };

  undoTask = (id) => {
    axios
      .put(endpoint + "/api/undoTask/" + id, {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
        },
      })
      .then((res) => {
        console.log(res);
        this.getTask();
      });
  };

  deleteTask = (id) => {
    axios
      .delete(endpoint + "/api/deleteTask/" + id, {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
        },
      })
      .then((res) => {
        console.log(res);
        this.getTask();
      });
  };

  render() {
    return (
      <div>
        <div className="row">
          <Header className="header" as="h2">
            <span
              style={{
                fontFamily: "cursive",
                fontSize: 50,
                fontFamily: "Brush Script MT, cursive",
              }}
            >
              TO-DO APP
            </span>
          </Header>
        </div>
        <div className="row">
          <Form
            onSubmit={this.onSubmit}
            style={{
              borderStyle: "outset",
              borderWidth: 4,
              borderColor: "grey",
            }}
          >
            <Input
              type="text"
              name="task"
              onChange={this.onChange}
              value={this.state.task}
              fluid
              placeholder="Create Task"
            />
            {/* <Button >Create Task</Button> */}
          </Form>
        </div>
        <div className="row">
          <Card.Group>{this.state.items}</Card.Group>
        </div>
      </div>
    );
  }
}

export default ToDo;
