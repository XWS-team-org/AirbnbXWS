import { CommonModule } from "@angular/common";
import { NgModule } from "@angular/core";
import { FormsModule } from "@angular/forms";
import { Routes, RouterModule } from "@angular/router";
import { AppRoutingModule } from "src/app/app-routing.module";
import { MaterialModule } from "src/app/material/material.module";
import { GuestProfileComponent } from "./guest-profile/guest-profile.component";
import { GuestEditProfileComponent } from "./guest-edit-profile/guest-edit-profile.component";
import { GuestHomeComponent } from "./guest-home/guest-home.component";
import { GuestToolbarComponent } from "./guest-toolbar/guest-toolbar.component";


const routes: Routes = [
    { path: 'guest-profile', component: GuestProfileComponent },
    { path: 'guest-edit', component: GuestEditProfileComponent},
    { path:'guest-home', component: GuestHomeComponent},
  ];
  
  
  
  @NgModule({
    declarations: [
        GuestEditProfileComponent,
        GuestProfileComponent,
        GuestHomeComponent,
        GuestToolbarComponent,
  ],
    imports: [
      CommonModule,
      AppRoutingModule,
      MaterialModule,
      FormsModule,
      RouterModule.forChild(routes)
    ]
  })
  export class GuestModule { }
  