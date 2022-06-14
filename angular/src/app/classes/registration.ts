export class Vehicle {
    type: string;
    make: string;
    model: string;
    colour: string;
}

export class Insurer {
    name: string;
    code: number;
}

export class Registration {
    number_plate: string;
    vehicle: Vehicle;
    insurer: Insurer;
}