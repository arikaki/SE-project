describe('The Profile Page', () => {
    it('successfully loads', () => {

      // cy.login();
      cy.visit('http://localhost:3000/profile') // change URL to match your dev URL
      cy.contains("Follow").should("exist");
      cy.contains("Answer").should("exist");
      cy.contains("Message").should("exist");
      cy.contains("Post").should("exist");
      cy.contains("Edit").should("exist");
    })

    it("routes to a correct pages", () => {
      cy.contains("HOME").click();
      cy.url().should("include", "/");
    });
  })