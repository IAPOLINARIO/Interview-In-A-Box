import React from "react";
import "./Header.css";
import { Avatar } from "@material-ui/core";
import SearchIcon from "@material-ui/icons/Search";
import { useStateValue } from "../../StateProvider";

function Header() {
  const [{ user }] = useStateValue();
  return (
    <div className="header">
      <div className="header-left">
        <Avatar
          className="header-avatar"
          alt={user?.displayName}
          src={user?.photoURL}
        />
      </div>
      <div className="header-search">
        <SearchIcon />
        <input placeholder="Search" type="text" />
      </div>
      <div className="header-right">
      </div>
    </div>
  );
}

export default Header;
