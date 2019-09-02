db.createUser(
    {
        user: "accountUser",
        pwd: "password",
        roles: [
            {
                role: "readWrite",
                db: "xw"
            }
        ]
    }
);