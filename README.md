# Power-4

## CSS structure

Styles were modularized to keep the design identical while making the code easier to navigate:

- src/CSS/base.css: shared base (body, buttons, header, common utilities)
- src/CSS/index.css: styles used by template/index.html
- src/CSS/diff.css: styles used by template/diff.html
- src/CSS/play.css: styles used by template/play.html
- src/CSS/regle.css: styles used by template/regle.html
- src/CSS/merch.css: styles used by template/merch.html (skins gallery + camera link)
- src/CSS/victoire.css: styles used by template/victoire.html
- src/CSS/egalite.css: styles used by template/egalite.html
- src/CSS/personalisation.css: styles used by template/personalisation.html
- src/CSS/camera.css: page-specific overrides for template/camera.html

Each HTML template now links to base.css plus its own page stylesheet. The original src/CSS/style.css remains in the repo for reference but is no longer used by templates.

Tip: Add new global utilities to base.css, and keep page-only rules in that page’s CSS to avoid regressions.

