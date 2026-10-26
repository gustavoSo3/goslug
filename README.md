GoSlug

TODO:

* [x] / gives you the welcome page
* [x] /admin takes you to slug admin page
* [x] /slug will redirect to slugs saved
* [x] /api
    * [x] POST /api/redirects will create a new redirect, no duplicate slugs
    * [x] DELETE /api/redirects will delete a redirect using the slug as reference

* [x] validate all fields - Current only if all aren't empty
* [ ] Slug must not contain spaces
* [ ] limit description to 100 chrs and slug to 20 and destination to max URL size

* [ ] Have persistan storage
    * [x] Database Interface
    * [x] In memory database
    * [ ] localfile storage
    * [ ] remote db (firestore/mongodb)

* [ ] User accounts? maybe just the admin, and on first set up must change??
    * Authentication should be a must before sharing or deployin more

* [ ] Better Loggin Of runn time logs

* [x] Created simple Wellcome page