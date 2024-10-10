db = db.getSiblingDB('tracker_service');

if (!db.getCollectionNames().includes('tracker')) {
    db.createCollection('tracker');
    db.tracker.insertMany([
        {"id":"e3651382-2f5b-4f17-afb5-83cbe919b38c","vehicle_id":"67890","location":[{"latitude":37.7749,"longitude":-122.4194},{"latitude":34.0522,"longitude":-118.2437}],"created_at":1633072800,"updated_at":1633072800}
    ]);
}
