import React from "react";
import HomeIcon from "@material-ui/icons/Home";
import FeaturedPlayListOutlinedIcon from "@material-ui/icons/FeaturedPlayListOutlined";
import {
  AssignmentTurnedInOutlined,
  NotificationsOutlined,
  PeopleAltOutlined,
  Search,
} from "@material-ui/icons";
import "./css/QuoraHeader.css";
import "react-responsive-modal/styles.css";
import { auth } from "../firebase";
import { signOut } from "firebase/auth";
import { logout, selectUser } from "../feature/userSlice";
import { useDispatch, useSelector } from "react-redux";
import { Navbar, Nav, Container, Button, Dropdown } from "react-bootstrap";
import NavDropdown from "react-bootstrap/NavDropdown";
// import DropdownButton from 'react-bootstrap/DropdownButton'

function QuoraHeader(props) {
  const dispatch = useDispatch();
  const user = useSelector(selectUser);

  const handleClick = () => {
    props.setShowFade(true);
  };

  const handleLogout = () => {
    if (window.confirm("Are you sure to logout ?")) {
      signOut(auth)
        .then(() => {
          dispatch(logout());
          console.log("Logged out");
        })
        .catch(() => {
          console.log("error in logout");
        });
    }
  };
  return (
    <div className="qHeader">
      <Navbar expand="lg">
        <Container fluid>
          <div className="qHeader__logo" style={{ paddingLeft: "20px" }}>
            <img src="/logo192.png" alt="logo" />
          </div>
          <Navbar.Toggle aria-controls="navbarScroll" />
          <Navbar.Collapse id="navbarScroll">
            <Nav
              className="mx-auto "
              style={{ maxHeight: "50px", paddingLeft: "30px" }}
              navbarScroll
            >
              <Nav.Link href="/">
                <HomeIcon />
              </Nav.Link>
              <Nav.Link href="/">
                <FeaturedPlayListOutlinedIcon />
              </Nav.Link>

              <Nav.Link href="/">
                <AssignmentTurnedInOutlined />
              </Nav.Link>
              <Nav.Link href="/">
                <PeopleAltOutlined />
              </Nav.Link>
              <Nav.Link href="/">
                <NotificationsOutlined />
              </Nav.Link>
            </Nav>

            <div className="qHeader__input">
              <Search />
              <input
                type="text"
                placeholder="Search questions"
                onClick={handleClick}
              />
            </div>
            <Nav
              className="me-auto"
              style={{
                position: "relative",
                left: "13rem",
              }}
            >
              <NavDropdown
                title={user?.userName.split(" ")[0]}
                style={{ backgroundColor: "#90EE90", borderRadius: "20px" }}
              >
                <NavDropdown.Item href="/profile">
                  <Button variant="primary">Profile</Button>
                </NavDropdown.Item>
                <NavDropdown.Item href="#action4">
                  <Button variant="danger" onClick={handleLogout}>
                    Log Out
                  </Button>
                </NavDropdown.Item>
              </NavDropdown>
            </Nav>
          </Navbar.Collapse>
        </Container>
      </Navbar>
    </div>
  );
}

export default QuoraHeader;
