import { Component } from '@angular/core';
import { registrationService } from './services/registration.service';

import { Registration } from './classes/registration';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'firt-angular';
  constructor(private registrationService: registrationService){}

  registrations: Registration[];

  ngOnInit() {
    this.registrationService.getRegistrations()
    .subscribe(
      data => {
        this.registrations = data;
      }
    )
  }

  createRegistration(data: any) {
    const body: Registration = {
      number_plate: data.value.number_plate,
      vehicle: {
        make: data.value.make,
        model: data.value.model,
        type: data.value.type,
        colour: data.value.colour,
      },
      insurer: {
        code: data.value.code,
        name: data.value.name
      }
    }
    this.registrationService.createRegistration(body).subscribe(() => {
      this.registrationService.getRegistrations().subscribe(data => {
        this.registrations = data
      })
    })
    data.reset()
  }

}
