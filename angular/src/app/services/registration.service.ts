import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { Registration } from "../classes/registration";
import { HttpClient } from "@angular/common/http";
@Injectable()
export class registrationService {
    constructor(private httpClient: HttpClient) {}
    getRegistrations(): Observable<any>{
        return this.httpClient.get("http://127.0.0.1:3000/registrations")
    }

    createRegistration(data: Registration) {
        return this.httpClient.post("http://127.0.0.1:3000/registration", data)
    }
}