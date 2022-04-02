describe('The Home Page', () => {
    it('successfully loads', () => {
      cy.visit('http://localhost:3000/profile   ') // change URL to match your dev URL
      cy.contains("Follow").should("exist");
      cy.contains("Contact").should("exist");
      cy.contains("Message").should("exist");
      cy.contains("Home").should("exist");
      cy.contains("Edit").should("exist");
    })
  })