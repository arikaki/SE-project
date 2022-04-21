describe('The Home Page', () => {
    it('successfully loads', () => {
      cy.visit('http://localhost:3000') // change URL to match your dev URL

      // cy.login();
      cy.contains("History").should("exist");
      cy.contains("Answers").should("exist");
      cy.contains("Business").should("exist");
      cy.contains("Discover").should("exist");
    })

    // it('successfully select dropdown', () => {
    //   cy.get('Post ').cy.get('select').select('Science').should('have.value', 'Science')
    // })

    it("routes to a correct pages", () => {
      cy.contains("Answers").click();
      cy.url().should("include", "answers");
    });
  })