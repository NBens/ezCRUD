# ezCRUD
A tool to generate CRUD pages using julienschmidt/httprouter, with support to multiple database systems &amp; templates

# ToDo:

* Suitable Template
* Suitable JSON Structure
* JSON Generator from a database
* Generate Config File (PHP config)
* Mix Templates With PHP Code

# JSON File Structure:
    {
        "database_type": "mysql",
        "database_host": "localhost",
        "database_user": "root",
        "database_pass": "root",
        "database_name": "MyCoolLilCrudApp",
        "tables": {
            "users": {
                "create": ["username", "password", "email"],
                "read": ["id", "username", "email"],
                "update": ["username", "password", "email"],
                "delete": ["id", "username"]
            },
            "categories": {
                "create": ["name", "description"],
                "read": ["id", "username"],
                "update": ["name", "description"],
                "delete": ["id", "name"]
            }
        }
    }
