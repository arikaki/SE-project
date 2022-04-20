import React from "react";
import "./css/NewCategories.css"

function NewCategories() {
  console.log("New Category rendered");
  return (
    <div>
      <div className="text-grid"> Select atleast five categories</div>
      <div className="container-grid">{/* <Grid /> */}</div>
    </div>
  );
}

export default NewCategories;
