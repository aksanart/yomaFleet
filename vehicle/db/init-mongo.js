db = db.getSiblingDB('vehicle_service');

if (!db.getCollectionNames().includes('vehicle')) {
    db.createCollection('vehicle');
    db.vehicle.insertMany([
        {"vehicle_name": 'name1', "vehicle_model": 'model1', "vehicle_status": 'idle', "license_number": 'license1', "mileage": 1, "id": 'e3651382-2f5b-4f17-afb5-83cbe919b38c',"created_at":1633072800,"updated_at":1633072800}
    ]);
}
