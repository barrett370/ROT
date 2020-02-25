import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { TopRoomsComponent } from './top-rooms.component';

describe('TopRoomsComponent', () => {
  let component: TopRoomsComponent;
  let fixture: ComponentFixture<TopRoomsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ TopRoomsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(TopRoomsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
