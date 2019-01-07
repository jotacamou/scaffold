Name:           scaffold
Version:        %{_version}
Release:        1
Summary:        Binary RPM package for sre/scaffold

License:        GPLv3
Source0:        ../SOURCES/%{name}
Source1:        ../SOURCES/config.yaml

%description
Scaffold helps create skeleton projects from templates.

%install
rm -rf %{buildroot}
install -d -m 0755 %{buildroot}/usr/local/bin
install -m 0755 %{SOURCE0} %{buildroot}/usr/local/bin
install -d -m 0755 %{buildroot}/opt/%{name}/etc
install -m 0644 %{SOURCE1} %{buildroot}/opt/%{name}/etc

%clean
rm -rf %{buildroot}

%files
/usr/local/bin/%{name}
/opt/%{name}/etc/config.yaml

%changelog
