describe('The Category Page', () => {
    it('successfully loads', () => {
      cy.visit('http://localhost:3000/categories') // change URL to match your dev URL

      // cy.login();
      cy.contains("Movies").should("exist");
      cy.contains("Technology").should("exist");
      cy.contains("Business").should("exist");
      cy.contains("Education").should("exist");
    })

    it("routes to a correct pages", () => {
      cy.contains("Continue").click();
      cy.url().should("include", "Continue");
    });
  })