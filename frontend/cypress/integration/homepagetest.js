describe('The Home Page', () => {
    it('successfully loads', () => {
      cy.visit('http://localhost:3000') // change URL to match your dev URL

      cy.contains("History").should("exist");
      cy.contains("Serach").should("exist");
      cy.contains("Business").should("exist");
      cy.contains("Discover").should("exist");
      cy.contains("question").should("exist");

    })

    it("routes to a correct pages", () => {
      cy.findAllByText("Answers").click();
      cy.url().should("include", "answers");
    });
  })