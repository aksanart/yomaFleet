db = db.getSiblingDB('api_gateway');

if (!db.getCollectionNames().includes('user')) {
    db.createCollection('user');
    db.user.insertMany([
        {"email":"admin@example.com","password":"$2a$10$kAgxNkgQG6xFYUMIOyD5geRr82YRHLg97BX6zt3K8N97Ii6RShemO","created_at":1633072800,"updated_at":1633072800}
    ]);
}
